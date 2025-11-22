package scaffold

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/RainDragonSk8er/AOC2025/pkg/aoc"
)

func TestRun(t *testing.T) {
	// Create a temporary directory for the test
	// This directory is automatically removed when the test finishes
	tempDir := t.TempDir()

	// Mock Leaderboard Data
	aliceName := "Alice"
	leaderboard := &aoc.JSONResponse{
		Members: map[string]aoc.Member{
			"1": {ID: 1, Name: &aliceName},
			"2": {ID: 2, Name: nil}, // Anonymous
		},
	}

	// Mock Mapping
	mapping := map[string]string{
		"Alice": "alice_wonderland",
	}

	// Run Scaffolder for Day 1
	err := Run(tempDir, 1, leaderboard, mapping)
	if err != nil {
		t.Fatalf("Run failed: %v", err)
	}

	// Verify Directories Created
	expectedDirs := []string{
		filepath.Join(tempDir, "solutions", "day01", "alice_wonderland"),
		filepath.Join(tempDir, "solutions", "day01", "Anonymous"),
	}

	for _, dir := range expectedDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			t.Errorf("Expected directory not created: %s", dir)
		}
	}

	// Verify Files Created (Default to .md)
	expectedFiles := []string{
		filepath.Join(tempDir, "solutions", "day01", "alice_wonderland", "main.md"),
		filepath.Join(tempDir, "solutions", "day01", "Anonymous", "main.md"),
	}

	for _, file := range expectedFiles {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			t.Errorf("Expected file not created: %s", file)
		}
	}

	// Test Smart Detection (Day 2)
	// Create a mock file for Day 1 to simulate previous language usage
	aliceDay1File := filepath.Join(tempDir, "solutions", "day01", "alice_wonderland", "main.py")
	os.WriteFile(aliceDay1File, []byte("print('hello')"), 0644)

	// Run Scaffolder for Day 2
	err = Run(tempDir, 2, leaderboard, mapping)
	if err != nil {
		t.Fatalf("Run day 2 failed: %v", err)
	}

	// Verify Day 2 file is main.py for Alice
	aliceDay2File := filepath.Join(tempDir, "solutions", "day02", "alice_wonderland", "main.py")
	if _, err := os.Stat(aliceDay2File); os.IsNotExist(err) {
		t.Errorf("Expected smart detection to create main.py, but file missing: %s", aliceDay2File)
	}
}
