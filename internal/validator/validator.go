package validator

import (
	"fmt"
	"strings"

	"learn-cli/internal/model"
	"learn-cli/internal/runner"
)

type TestResult struct {
	Name     string
	Hidden   bool
	Passed   bool
	ErrorMsg string
	Diff     string
}

// Normalize cleans up whitespace, CRLF, and trailing spaces per line for fair comparisons.
func Normalize(s string) string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	lines := strings.Split(s, "\n")
	var cleaned []string
	for _, line := range lines {
		cleaned = append(cleaned, strings.TrimRight(line, " \t\r"))
	}
	result := strings.Join(cleaned, "\n")
	return strings.TrimSpace(result)
}

// Evaluate checks student run result against the test case specifications.
func Evaluate(tc model.TestCase, res *runner.RunResult) *TestResult {
	if res.TimedOut {
		return &TestResult{
			Name:     tc.Name,
			Hidden:   tc.Hidden,
			Passed:   false,
			ErrorMsg: "[✗] TIMEOUT (> 3 detik). Periksa perulangan while/for kamu, kemungkinan terjadi infinite loop!",
		}
	}

	normExpected := Normalize(tc.Expected)
	normActual := Normalize(res.Stdout)

	if normActual == normExpected {
		return &TestResult{
			Name:   tc.Name,
			Hidden: tc.Hidden,
			Passed: true,
		}
	}

	// Mismatch occurred
	var diff string
	if tc.Hidden {
		diff = fmt.Sprintf("Uji kasus tersembunyi gagal pada skenario: %s\n(Detail output disembunyikan agar kamu mencoba memperbaikinya sendiri)", tc.Name)
	} else {
		diff = fmt.Sprintf("--- Input Pengujian:\n%s\n--- Output Diharapkan:\n%s\n--- Output Diterima:\n%s",
			strings.TrimSpace(tc.Stdin),
			normExpected,
			normActual,
		)
	}

	errMsg := ""
	if res.ExitCode != 0 {
		errMsg = fmt.Sprintf("Program berhenti dengan kode kesalahan (Exit Code: %d). Stderr: %s", res.ExitCode, res.Stderr)
	}

	return &TestResult{
		Name:     tc.Name,
		Hidden:   tc.Hidden,
		Passed:   false,
		ErrorMsg: errMsg,
		Diff:     diff,
	}
}
