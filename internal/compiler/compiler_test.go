package compiler

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCompilerCheckAndCompile(t *testing.T) {
	_, err := CheckCompiler("c")
	if err != nil {
		t.Fatalf("expected gcc to be installed: %v", err)
	}

	tempDir, err := os.MkdirTemp("", "learn-comp-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	src := filepath.Join(tempDir, "main.c")
	bin := filepath.Join(tempDir, "main.exe")
	os.WriteFile(src, []byte("#include <stdio.h>\nint main(){ printf(\"OK\"); return 0; }"), 0644)

	output, err := Compile(src, bin)
	if err != nil {
		t.Fatalf("compile failed: %v, output: %s", err, output)
	}

	if _, err := os.Stat(bin); os.IsNotExist(err) {
		t.Errorf("binary was not created")
	}
}
