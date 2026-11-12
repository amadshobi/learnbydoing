package scaffold

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInitStageSafeOverwrite(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "learn-scaffold-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// First init
	res, err := InitStage(tempDir, "easy", 1, false)
	if err != nil {
		t.Fatalf("first init failed: %v", err)
	}

	if res.FolderName != "easy-stage-1" {
		t.Errorf("expected folder name easy-stage-1, got %s", res.FolderName)
	}

	// Modify main.c
	stageDir := filepath.Join(tempDir, "learn-c", res.FolderName)
	mainC := filepath.Join(stageDir, "main.c")
	err = os.WriteFile(mainC, []byte("// User modification"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Second init without force must fail to protect user code
	_, err = InitStage(tempDir, "easy", 1, false)
	if err == nil {
		t.Fatalf("expected error on re-init without force, but got nil")
	}

	// Second init with force must succeed
	_, err = InitStage(tempDir, "easy", 1, true)
	if err != nil {
		t.Fatalf("expected init with force to succeed, got %v", err)
	}

	// Verify main.c was overwritten
	content, err := os.ReadFile(mainC)
	if err != nil {
		t.Fatal(err)
	}
	if string(content) == "// User modification" {
		t.Errorf("expected main.c to be overwritten when force is true")
	}
}
