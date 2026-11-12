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

func FindProgressFilePath(startDir string) string {
	// Check if learn-c exists in startDir
	learnCDir := filepath.Join(startDir, "learn-c")
	if info, err := os.Stat(learnCDir); err == nil && info.IsDir() {
		return filepath.Join(learnCDir, ProgressFileName)
	}

	curr := startDir
	for i := 0; i < 4; i++ {
		candidate := filepath.Join(curr, ProgressFileName)
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
		parent := filepath.Dir(curr)
		if parent == curr {
			break
		}
		curr = parent
	}
	return filepath.Join(startDir, ProgressFileName)
}

func LoadProgress(rootDir string) (*ProgressData, error) {
	filePath := FindProgressFilePath(rootDir)
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

	filePath := FindProgressFilePath(rootDir)
	return os.WriteFile(filePath, data, 0644)
}
