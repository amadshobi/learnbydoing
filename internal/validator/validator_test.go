package validator

import (
	"strings"
	"testing"

	"learn-cli/internal/model"
	"learn-cli/internal/runner"
)

func TestEvaluateNormalizedCRLF(t *testing.T) {
	tc := model.TestCase{
		Name:     "Test Penjumlahan",
		Hidden:   false,
		Stdin:    "5\n3\n",
		Expected: "Hasil: 8\n",
	}
	runRes := &runner.RunResult{
		Stdout:   "Hasil: 8\r\n", // Windows CRLF
		ExitCode: 0,
		TimedOut: false,
	}
	res := Evaluate(tc, runRes)
	if !res.Passed {
		t.Errorf("expected test to pass despite CRLF differences, got diff: %s", res.Diff)
	}
}

func TestEvaluateMismatchDiff(t *testing.T) {
	tc := model.TestCase{
		Name:     "Test Perkalian",
		Hidden:   false,
		Stdin:    "4\n5\n",
		Expected: "Hasil: 20\n",
	}
	runRes := &runner.RunResult{
		Stdout:   "Hasil: 10\n",
		ExitCode: 0,
		TimedOut: false,
	}
	res := Evaluate(tc, runRes)
	if res.Passed {
		t.Errorf("expected test to fail due to mismatch")
	}
	if !strings.Contains(res.Diff, "Diharapkan") || !strings.Contains(res.Diff, "Diterima") {
		t.Errorf("expected diff to show expected and received output, got: %s", res.Diff)
	}
}

func TestEvaluateHiddenTestFailure(t *testing.T) {
	tc := model.TestCase{
		Name:     "Skenario Pembagian Nol (Hidden)",
		Hidden:   true,
		Stdin:    "10\n0\n",
		Expected: "Error: Tidak bisa membagi dengan nol!\n",
	}
	runRes := &runner.RunResult{
		Stdout:   "Hasil: 0\n",
		ExitCode: 0,
		TimedOut: false,
	}
	res := Evaluate(tc, runRes)
	if res.Passed {
		t.Errorf("expected hidden test to fail")
	}
	// For hidden tests, diff should NOT leak raw expected answer
	if strings.Contains(res.Diff, "Error: Tidak bisa membagi dengan nol!") {
		t.Errorf("diff should not leak raw expected answer for hidden tests")
	}
}

func TestEvaluateTimeout(t *testing.T) {
	tc := model.TestCase{
		Name:     "Test Loop",
		Hidden:   false,
		Stdin:    "",
		Expected: "Selesai\n",
	}
	runRes := &runner.RunResult{
		TimedOut: true,
	}
	res := Evaluate(tc, runRes)
	if res.Passed {
		t.Errorf("expected timeout test to fail")
	}
	if !strings.Contains(res.ErrorMsg, "TIMEOUT") {
		t.Errorf("expected error message to mention TIMEOUT, got: %s", res.ErrorMsg)
	}
}
