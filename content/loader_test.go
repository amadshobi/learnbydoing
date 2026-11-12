package content

import (
	"testing"
)

func TestGetStageContent(t *testing.T) {
	c, err := GetStageContent("easy", 1)
	if err != nil {
		t.Fatalf("failed to get stage 1 content: %v", err)
	}
	if len(c.Tests) < 2 {
		t.Errorf("expected at least 2 tests (sample and hidden), got %d", len(c.Tests))
	}
	if len(c.StarterCode) == 0 {
		t.Errorf("expected starter code to be present")
	}
	if len(c.Readme) == 0 {
		t.Errorf("expected readme to be present")
	}

	// Test non-existent stage
	_, err = GetStageContent("easy", 99)
	if err == nil {
		t.Errorf("expected error for non-existent stage, got nil")
	}
}
