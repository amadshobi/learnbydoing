package updater

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const RepoOwner = "amadshobi"
const RepoName = "learnbydoing"

type GitHubRelease struct {
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

func CheckUpdate(currentVersion string) (*GitHubRelease, bool, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", RepoOwner, RepoName)
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, false, err
	}
	req.Header.Set("User-Agent", "learnbydoing-cli")

	resp, err := client.Do(req)
	if err != nil {
		return nil, false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, false, fmt.Errorf("GitHub API mengembalikan status %s", resp.Status)
	}

	var rel GitHubRelease
	if err := json.NewDecoder(resp.Body).Decode(&rel); err != nil {
		return nil, false, err
	}

	hasUpdate := rel.TagName != "" && rel.TagName != currentVersion
	return &rel, hasUpdate, nil
}

func PerformUpdate(rel *GitHubRelease, out io.Writer) error {
	assetName := fmt.Sprintf("learn-%s-%s", runtime.GOOS, runtime.GOARCH)
	if runtime.GOOS == "windows" {
		assetName += ".exe"
	}

	var downloadURL string
	for _, a := range rel.Assets {
		if a.Name == assetName {
			downloadURL = a.BrowserDownloadURL
			break
		}
	}

	if downloadURL == "" {
		return fmt.Errorf("binary untuk platform %s/%s (%s) tidak ditemukan pada rilis %s", runtime.GOOS, runtime.GOARCH, assetName, rel.TagName)
	}

	fmt.Fprintf(out, "⬇️  Mengunduh versi terbaru %s (%s)...\n", rel.TagName, assetName)

	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Get(downloadURL)
	if err != nil {
		return fmt.Errorf("gagal mengunduh: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download gagal dengan status: %s", resp.Status)
	}

	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("gagal menemukan lokasi executable saat ini: %w", err)
	}
	exePath, err = filepath.EvalSymlinks(exePath)
	if err != nil {
		return fmt.Errorf("gagal resolve symlink executable: %w", err)
	}

	tempFile := exePath + ".new"
	f, err := os.OpenFile(tempFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return fmt.Errorf("gagal membuat file temporary: %w", err)
	}
	_, err = io.Copy(f, resp.Body)
	f.Close()
	if err != nil {
		_ = os.Remove(tempFile)
		return fmt.Errorf("gagal menulis binary: %w", err)
	}

	oldFile := exePath + ".old"
	_ = os.Remove(oldFile)

	if err := os.Rename(exePath, oldFile); err != nil {
		_ = os.Remove(tempFile)
		return fmt.Errorf("gagal backup binary lama: %w", err)
	}

	if err := os.Rename(tempFile, exePath); err != nil {
		_ = os.Rename(oldFile, exePath)
		return fmt.Errorf("gagal memasang binary baru: %w", err)
	}

	_ = os.Remove(oldFile)
	return nil
}
