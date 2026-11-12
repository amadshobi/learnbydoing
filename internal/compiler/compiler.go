package compiler

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// CheckCompiler checks if the necessary compiler (gcc or g++) is available in PATH.
func CheckCompiler(lang string) (string, error) {
	cmdName := "gcc"
	if strings.ToLower(lang) == "cpp" || strings.ToLower(lang) == "c++" {
		cmdName = "g++"
	}

	path, err := exec.LookPath(cmdName)
	if err != nil {
		return "", fmt.Errorf("compiler %q tidak ditemukan di sistem Anda.\nSilakan install MinGW-w64 (Windows) atau build-essential (Linux) dan pastikan %s sudah masuk ke PATH", cmdName, cmdName)
	}
	return path, nil
}

// Compile compiles a C or C++ source file into an executable binary.
func Compile(sourcePath string, outputPath string) (string, error) {
	ext := strings.ToLower(filepath.Ext(sourcePath))
	var cmd *exec.Cmd

	if ext == ".cpp" || ext == ".cc" {
		if _, err := CheckCompiler("cpp"); err != nil {
			return "", err
		}
		cmd = exec.Command("g++", "-std=c++14", sourcePath, "-o", outputPath)
	} else {
		if _, err := CheckCompiler("c"); err != nil {
			return "", err
		}
		cmd = exec.Command("gcc", "-std=c11", sourcePath, "-o", outputPath, "-lm")
	}

	outputBytes, err := cmd.CombinedOutput()
	outputStr := string(outputBytes)
	if err != nil {
		return outputStr, fmt.Errorf("kompilasi gagal: %w\n%s", err, outputStr)
	}

	return outputStr, nil
}
