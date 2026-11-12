package runner

import (
	"bytes"
	"context"
	"os/exec"
	"time"
)

type RunResult struct {
	Stdout   string
	Stderr   string
	ExitCode int
	TimedOut bool
}

// RunBinary runs a compiled executable binary with stdin input and a timeout.
func RunBinary(binaryPath string, stdinInput string, timeout time.Duration) (*RunResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, binaryPath)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	if len(stdinInput) > 0 {
		cmd.Stdin = bytes.NewBufferString(stdinInput)
	}

	err := cmd.Run()

	res := &RunResult{
		Stdout: stdoutBuf.String(),
		Stderr: stderrBuf.String(),
	}

	if ctx.Err() == context.DeadlineExceeded {
		res.TimedOut = true
		return res, nil
	}

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			res.ExitCode = exitErr.ExitCode()
			return res, nil
		}
		return res, err
	}

	res.ExitCode = 0
	return res, nil
}
