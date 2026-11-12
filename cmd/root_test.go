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

func TestInitForceFlagAfterPositionalArg(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "learn-cmd-force-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	origWd, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(origWd)

	// First init
	var outBuf bytes.Buffer
	_ = ExecuteArgs([]string{"init", "1"}, &outBuf)

	// Re-run with --force placed AFTER the stage number (documented usage)
	outBuf.Reset()
	err = ExecuteArgs([]string{"init", "1", "--force"}, &outBuf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	outStr := outBuf.String()
	if !strings.Contains(outStr, "berhasil dibuat") {
		t.Errorf("expected --force after stage number to succeed, but got: %s", outStr)
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
	if !strings.Contains(outStr, "STAGED LEARNING") {
		t.Errorf("expected help message to contain STAGED LEARNING banner, got: %s", outStr)
	}
	if !strings.Contains(outStr, "RESET aman") {
		t.Errorf("expected help message to explain --force as safe reset, got: %s", outStr)
	}
}

func TestStageListCommand(t *testing.T) {
	var outBuf bytes.Buffer
	err := ExecuteArgs([]string{"stage", "list"}, &outBuf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	outStr := outBuf.String()
	if !strings.Contains(outStr, "DAFTAR STAGE TANTANGAN") {
		t.Errorf("expected stage list header, got: %s", outStr)
	}
	if !strings.Contains(outStr, "1. Kalkulator Sederhana") || !strings.Contains(outStr, "learn init 1 --easy") {
		t.Errorf("expected stage 1 in stage list, got: %s", outStr)
	}
	if !strings.Contains(outStr, "2. Tebak Angka") {
		t.Errorf("expected stage 2 in stage list, got: %s", outStr)
	}
	if !strings.Contains(outStr, "3. Kasir Mini") {
		t.Errorf("expected stage 3 in stage list, got: %s", outStr)
	}

	// Test alias 'learn list'
	outBuf.Reset()
	err = ExecuteArgs([]string{"list"}, &outBuf)
	if err != nil {
		t.Fatalf("unexpected error with 'learn list': %v", err)
	}
	if !strings.Contains(outBuf.String(), "DAFTAR STAGE TANTANGAN") {
		t.Errorf("expected 'learn list' to work, got: %s", outBuf.String())
	}
}

func TestVersionCommand(t *testing.T) {
	var outBuf bytes.Buffer
	err := ExecuteArgs([]string{"version"}, &outBuf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	outStr := outBuf.String()
	if !strings.Contains(outStr, "learnbydoing CLI v") {
		t.Errorf("expected version output, got: %s", outStr)
	}

	outBuf.Reset()
	err = ExecuteArgs([]string{"-v"}, &outBuf)
	if err != nil {
		t.Fatalf("unexpected error with -v: %v", err)
	}
	if !strings.Contains(outBuf.String(), "learnbydoing CLI v") {
		t.Errorf("expected version output with -v, got: %s", outBuf.String())
	}
}

func TestUpdateCommandLatest(t *testing.T) {
	var outBuf bytes.Buffer
	err := ExecuteArgs([]string{"update"}, &outBuf)
	if err != nil {
		t.Fatalf("unexpected error on update: %v", err)
	}
	outStr := outBuf.String()
	if !strings.Contains(outStr, "Memeriksa pembaruan") {
		t.Errorf("expected update to check github, got: %s", outStr)
	}
}
