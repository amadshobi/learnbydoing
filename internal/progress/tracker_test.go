package progress

import (
	"os"
	"testing"
)

func TestProgressSaveAndLoad(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "learn-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	err = SaveProgress(tempDir, "easy", 1)
	if err != nil {
		t.Fatalf("unexpected error saving: %v", err)
	}

	p, err := LoadProgress(tempDir)
	if err != nil {
		t.Fatalf("unexpected error loading: %v", err)
	}
	if p.Completed["easy-1"] != true {
		t.Errorf("expected easy-1 to be completed")
	}
}
