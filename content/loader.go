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

	titleParts := strings.SplitN(stageDirName, "-", 2)
	title := stageDirName
	if len(titleParts) == 2 {
		title = strings.ReplaceAll(titleParts[1], "-", " ")
	}

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
