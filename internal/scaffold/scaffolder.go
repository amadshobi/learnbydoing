package scaffold

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"learn-cli/content"
	"learn-cli/internal/model"
)

type InitResult struct {
	StageDir   string
	FolderName string
	Title      string
}

const WorkspaceFolderName = "learn-c"
const StageMetaFileName = ".stage.json"

// InitStage prepares the folder structure for a given challenge track and stage.
func InitStage(targetRoot string, track string, stage int, force bool) (*InitResult, error) {
	stageContent, err := content.GetStageContent(track, stage)
	if err != nil {
		return nil, err
	}

	// Base workspace path
	baseDir := filepath.Join(targetRoot, WorkspaceFolderName)
	if filepath.Base(targetRoot) == WorkspaceFolderName {
		baseDir = targetRoot
	}

	stageDir := filepath.Join(baseDir, stageContent.FolderName)
	if err := os.MkdirAll(stageDir, 0755); err != nil {
		return nil, fmt.Errorf("gagal membuat folder stage: %w", err)
	}

	mainCPath := filepath.Join(stageDir, "main.c")
	mainCppPath := filepath.Join(stageDir, "main.cpp")

	// Check if solution file already exists and force is false
	hasExistingFile := false
	var existingName string
	if _, err := os.Stat(mainCPath); err == nil {
		hasExistingFile = true
		existingName = "main.c"
	} else if _, err := os.Stat(mainCppPath); err == nil {
		hasExistingFile = true
		existingName = "main.cpp"
	}

	if hasExistingFile && !force {
		return nil, fmt.Errorf("file kode %q sudah ada di %s.\nGunakan flag --force jika ingin menimpa (overwrite) file yang ada", existingName, stageDir)
	}

	// Write README.md
	readmePath := filepath.Join(stageDir, "README.md")
	if err := os.WriteFile(readmePath, []byte(stageContent.Readme), 0644); err != nil {
		return nil, fmt.Errorf("gagal menulis README.md: %w", err)
	}

	// Write starter code
	if err := os.WriteFile(mainCPath, []byte(stageContent.StarterCode), 0644); err != nil {
		return nil, fmt.Errorf("gagal menulis starter main.c: %w", err)
	}

	// Write .stage.json metadata
	meta := model.StageMeta{
		Track: track,
		Stage: stage,
		Title: stageContent.Title,
	}
	metaBytes, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("gagal serialisasi metadata: %w", err)
	}

	metaPath := filepath.Join(stageDir, StageMetaFileName)
	if err := os.WriteFile(metaPath, metaBytes, 0644); err != nil {
		return nil, fmt.Errorf("gagal menulis .stage.json: %w", err)
	}

	return &InitResult{
		StageDir:   stageDir,
		FolderName: stageContent.FolderName,
		Title:      stageContent.Title,
	}, nil
}
