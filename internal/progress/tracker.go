package progress

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const ProgressFileName = ".learn_progress.json"

type ProgressData struct {
	Completed map[string]bool `json:"completed"`
}

func getProgressFilePath(rootDir string) string {
	return filepath.Join(rootDir, ProgressFileName)
}

func LoadProgress(rootDir string) (*ProgressData, error) {
	filePath := getProgressFilePath(rootDir)
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return &ProgressData{Completed: make(map[string]bool)}, nil
		}
		return nil, err
	}

	var p ProgressData
	if err := json.Unmarshal(data, &p); err != nil {
		return &ProgressData{Completed: make(map[string]bool)}, nil
	}
	if p.Completed == nil {
		p.Completed = make(map[string]bool)
	}
	return &p, nil
}

func SaveProgress(rootDir string, track string, stage int) error {
	p, err := LoadProgress(rootDir)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("%s-%d", track, stage)
	p.Completed[key] = true

	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}

	filePath := getProgressFilePath(rootDir)
	return os.WriteFile(filePath, data, 0644)
}
