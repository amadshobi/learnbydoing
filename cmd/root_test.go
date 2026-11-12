package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestInitCommandExecution(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "learn-cmd-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	origWd, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(origWd)

	var outBuf bytes.Buffer
	err = ExecuteArgs([]string{"init", "1", "--easy"}, &outBuf)
	if err != nil {
		t.Fatalf("unexpected error executing init: %v", err)
	}

	outStr := outBuf.String()
	if !strings.Contains(outStr, "berhasil dibuat") {
		t.Errorf("expected success message in output, got: %s", outStr)
	}

	stageDir := filepath.Join(tempDir, "learn-c", "easy-stage-1")
	if _, err := os.Stat(filepath.Join(stageDir, "main.c")); os.IsNotExist(err) {
		t.Errorf("main.c was not created")
	}
}

func TestHelpCommandOutput(t *testing.T) {
	var outBuf bytes.Buffer
	err := ExecuteArgs([]string{"--help"}, &outBuf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	outStr := outBuf.String()
	if !strings.Contains(outStr, "learn init") || !strings.Contains(outStr, "learn test") {
		t.Errorf("expected help message to list commands, got: %s", outStr)
	}
}
