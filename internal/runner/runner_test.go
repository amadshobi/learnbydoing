package runner

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"learn-cli/internal/compiler"
)

func TestCompileAndRunEcho(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "learn-runner-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	cCode := `#include <stdio.h>
int main() {
    int a;
    if (scanf("%d", &a) == 1) {
        printf("Angka: %d\n", a * 2);
    }
    return 0;
}`
	src := filepath.Join(tempDir, "main.c")
	os.WriteFile(src, []byte(cCode), 0644)
	bin := filepath.Join(tempDir, "solution.exe")

	warns, err := compiler.Compile(src, bin)
	if err != nil {
		t.Fatalf("compile failed: %v, output: %s", err, warns)
	}

	res, err := RunBinary(bin, "21\n", 3*time.Second)
	if err != nil {
		t.Fatalf("run failed: %v", err)
	}
	if res.TimedOut {
		t.Errorf("did not expect timeout")
	}
	expected := "Angka: 42\n"
	actual := strings.ReplaceAll(res.Stdout, "\r\n", "\n")
	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestInfiniteLoopTimeout(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "learn-runner-timeout-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	cCode := `#include <stdio.h>
int main() {
    while(1) {}
    return 0;
}`
	src := filepath.Join(tempDir, "main.c")
	os.WriteFile(src, []byte(cCode), 0644)
	bin := filepath.Join(tempDir, "infinite.exe")

	_, err = compiler.Compile(src, bin)
	if err != nil {
		t.Fatalf("compile failed: %v", err)
	}

	res, err := RunBinary(bin, "", 1*time.Second)
	if err != nil {
		t.Fatalf("run failed: %v", err)
	}
	if !res.TimedOut {
		t.Errorf("expected process to timeout and be killed")
	}
}
