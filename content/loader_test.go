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

func TestListStages(t *testing.T) {
	stages, err := ListStages("easy")
	if err != nil {
		t.Fatalf("unexpected error listing stages: %v", err)
	}
	if len(stages) != 3 {
		t.Fatalf("expected 3 stages in easy track, got %d", len(stages))
	}
	if stages[0].Stage != 1 || stages[0].Title != "Kalkulator Sederhana" {
		t.Errorf("expected stage 1 Kalkulator Sederhana, got %d %s", stages[0].Stage, stages[0].Title)
	}
	if stages[1].Stage != 2 || stages[1].Title != "Tebak Angka" {
		t.Errorf("expected stage 2 Tebak Angka, got %d %s", stages[1].Stage, stages[1].Title)
	}
	if stages[2].Stage != 3 || stages[2].Title != "Kasir Mini" {
		t.Errorf("expected stage 3 Kasir Mini, got %d %s", stages[2].Stage, stages[2].Title)
	}
}
