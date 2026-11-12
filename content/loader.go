package content

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"path"
	"strings"

	"learn-cli/internal/model"
)

//go:embed easy/*
var embeddedFiles embed.FS

type StageContent struct {
	Track       string           `json:"track"`
	Stage       int              `json:"stage"`
	Title       string           `json:"title"`
	FolderName  string           `json:"folder_name"`
	Readme      string           `json:"readme"`
	StarterCode string           `json:"starter_code"`
	Tests       []model.TestCase `json:"tests"`
}

type StageInfo struct {
	Track string `json:"track"`
	Stage int    `json:"stage"`
	Title string `json:"title"`
}

func extractTitleFromReadme(readme string, fallback string) string {
	lines := strings.Split(readme, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "# ") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}
			return strings.TrimSpace(strings.TrimPrefix(line, "# "))
		}
	}
	return fallback
}

func ListStages(track string) ([]StageInfo, error) {
	trackDir := track
	entries, err := fs.ReadDir(embeddedFiles, trackDir)
	if err != nil {
		return nil, fmt.Errorf("track %q tidak ditemukan", track)
	}

	var stages []StageInfo
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		parts := strings.SplitN(entry.Name(), "-", 2)
		if len(parts) < 2 {
			continue
		}
		var stageNum int
		if _, err := fmt.Sscanf(parts[0], "%d", &stageNum); err != nil {
			continue
		}

		stagePath := path.Join(trackDir, entry.Name())
		readmeBytes, err := embeddedFiles.ReadFile(path.Join(stagePath, "README.md"))
		fallbackTitle := strings.ReplaceAll(parts[1], "-", " ")
		title := fallbackTitle
		if err == nil {
			title = extractTitleFromReadme(string(readmeBytes), fallbackTitle)
		}

		stages = append(stages, StageInfo{
			Track: track,
			Stage: stageNum,
			Title: title,
		})
	}

	// Sort stages by Stage ascending
	for i := 0; i < len(stages); i++ {
		for j := i + 1; j < len(stages); j++ {
			if stages[i].Stage > stages[j].Stage {
				stages[i], stages[j] = stages[j], stages[i]
			}
		}
	}

	return stages, nil
}

func GetStageContent(track string, stage int) (*StageContent, error) {
	trackDir := track
	entries, err := fs.ReadDir(embeddedFiles, trackDir)
	if err != nil {
		return nil, fmt.Errorf("track %q tidak ditemukan", track)
	}

	prefix := fmt.Sprintf("%d-", stage)
	var stageDirName string
	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), prefix) {
			stageDirName = entry.Name()
			break
		}
	}

	if stageDirName == "" {
		return nil, fmt.Errorf("stage %d pada track %s tidak ditemukan", stage, track)
	}

	stagePath := path.Join(trackDir, stageDirName)

	readmeBytes, err := embeddedFiles.ReadFile(path.Join(stagePath, "README.md"))
	if err != nil {
		return nil, fmt.Errorf("gagal membaca README.md: %w", err)
	}

	starterBytes, err := embeddedFiles.ReadFile(path.Join(stagePath, "starter.c"))
	if err != nil {
		return nil, fmt.Errorf("gagal membaca starter.c: %w", err)
	}

	testsBytes, err := embeddedFiles.ReadFile(path.Join(stagePath, "tests.json"))
	if err != nil {
		return nil, fmt.Errorf("gagal membaca tests.json: %w", err)
	}

	var tests []model.TestCase
	if err := json.Unmarshal(testsBytes, &tests); err != nil {
		return nil, fmt.Errorf("gagal parsing tests.json: %w", err)
	}

	fallbackTitle := stageDirName
	titleParts := strings.SplitN(stageDirName, "-", 2)
	if len(titleParts) == 2 {
		fallbackTitle = strings.ReplaceAll(titleParts[1], "-", " ")
	}
	title := extractTitleFromReadme(string(readmeBytes), fallbackTitle)

	return &StageContent{
		Track:       track,
		Stage:       stage,
		Title:       title,
		FolderName:  fmt.Sprintf("%s-stage-%d", track, stage),
		Readme:      string(readmeBytes),
		StarterCode: string(starterBytes),
		Tests:       tests,
	}, nil
}
