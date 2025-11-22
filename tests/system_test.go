package tests

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/RainDragonSk8er/AOC2025/pkg/aoc"
	"github.com/RainDragonSk8er/AOC2025/pkg/scaffold"
	"github.com/RainDragonSk8er/AOC2025/pkg/table"
)

func TestFullScaleSystem(t *testing.T) {
	// 1. Setup Temporary Environment
	tempDir := t.TempDir()

	// 2. Define Mock Data Scenarios
	aliceName := "Alice"
	bobName := "Bob"
	charlieName := "Charlie"

	scenarios := []struct {
		name        string
		leaderboard *aoc.JSONResponse
		date        time.Time
		expectTable bool
		expectText  []string // Text that MUST be present
	}{
		{
			name: "November 30 - Hidden",
			leaderboard: &aoc.JSONResponse{
				Members: map[string]aoc.Member{
					"1": {ID: 1, Name: &aliceName, Stars: 50, LocalScore: 100},
				},
			},
			date:        time.Date(2025, time.November, 30, 23, 59, 0, 0, time.UTC),
			expectTable: false,
			expectText:  []string{"Leaderboard will be revealed"},
		},
		{
			name: "December 1 - Shown - Low Stars",
			leaderboard: &aoc.JSONResponse{
				Members: map[string]aoc.Member{
					"1": {ID: 1, Name: &aliceName, Stars: 2, LocalScore: 10},
				},
			},
			date:        time.Date(2025, time.December, 1, 12, 0, 0, 0, time.UTC),
			expectTable: true,
			expectText: []string{
				"Alice",
				"2 Stars",
				"(╯°□°)╯︵ ┻━┻", // Default low star emoticon
			},
		},
		{
			name: "December 12 - High Stars - Custom Theme",
			leaderboard: &aoc.JSONResponse{
				Members: map[string]aoc.Member{
					"1": {ID: 1, Name: &aliceName, Stars: 24, LocalScore: 500}, // Max stars
					"2": {ID: 2, Name: &bobName, Stars: 10, LocalScore: 200},
					"3": {ID: 3, Name: &charlieName, Stars: 0, LocalScore: 0},
				},
			},
			date:        time.Date(2025, time.December, 12, 12, 0, 0, 0, time.UTC),
			expectTable: true,
			expectText: []string{
				"Alice", "24 Stars", "(⌐■_■) 🎄", // Assuming default theme or we pass one
				"Bob", "10 Stars", "(._.)",
				"Charlie", "0 Stars", "(╯°□°)╯︵ ┻━┻",
			},
		},
	}

	// Load Default Theme for testing
	theme := table.Theme{
		Bar: table.BarConfig{Filled: "#", Empty: "."},
		Emoticons: []table.EmoticonConfig{
			{Threshold: 0, Icon: "(╯°□°)╯︵ ┻━┻"},
			{Threshold: 5, Icon: "(ಠ_ಠ)"},
			{Threshold: 9, Icon: "(._.)"},
			{Threshold: 13, Icon: "(•_•)"},
			{Threshold: 17, Icon: "(＾▽＾)"},
			{Threshold: 21, Icon: "(⌐■_■) 🎄"},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.name, func(t *testing.T) {
			output := table.Generate(sc.leaderboard, theme, sc.date)

			if sc.expectTable {
				if strings.Contains(output, "Leaderboard will be revealed") {
					t.Errorf("Expected table, got hidden message")
				}
			} else {
				if !strings.Contains(output, "Leaderboard will be revealed") {
					t.Errorf("Expected hidden message, got table")
				}
			}

			for _, txt := range sc.expectText {
				if !strings.Contains(output, txt) {
					t.Errorf("Expected output to contain %q, but it didn't.\nOutput:\n%s", txt, output)
				}
			}
		})
	}

	// 3. Test Scaffolding & Language Detection
	t.Run("Scaffolding Integration", func(t *testing.T) {
		// Setup: Alice has Day 1 Python solution
		day1Dir := filepath.Join(tempDir, "solutions", "day01", "Alice")
		if err := os.MkdirAll(day1Dir, 0755); err != nil {
			t.Fatal(err)
		}
		os.WriteFile(filepath.Join(day1Dir, "main.py"), []byte("print('prev')"), 0644)

		// Setup: Bob has Day 1 Rust solution
		day1Bob := filepath.Join(tempDir, "solutions", "day01", "Bob")
		if err := os.MkdirAll(day1Bob, 0755); err != nil {
			t.Fatal(err)
		}
		os.WriteFile(filepath.Join(day1Bob, "main.rs"), []byte("fn main() {}"), 0644)

		// Mock Leaderboard for Day 2
		lb := &aoc.JSONResponse{
			Members: map[string]aoc.Member{
				"1": {ID: 1, Name: &aliceName},
				"2": {ID: 2, Name: &bobName},
				"3": {ID: 3, Name: &charlieName}, // New user
			},
		}
		mapping := map[string]string{} // No mapping

		// Run Scaffolder for Day 2
		err := scaffold.Run(tempDir, 2, lb, mapping)
		if err != nil {
			t.Fatalf("Scaffold run failed: %v", err)
		}

		// Verify Alice -> main.py
		aliceFile := filepath.Join(tempDir, "solutions", "day02", "Alice", "main.py")
		if _, err := os.Stat(aliceFile); os.IsNotExist(err) {
			t.Errorf("Alice should have main.py")
		}
		content, _ := os.ReadFile(aliceFile)
		if !strings.Contains(string(content), "I'm not slow") {
			t.Errorf("Alice's file should contain Python joke")
		}

		// Verify Bob -> main.rs
		bobFile := filepath.Join(tempDir, "solutions", "day02", "Bob", "main.rs")
		if _, err := os.Stat(bobFile); os.IsNotExist(err) {
			t.Errorf("Bob should have main.rs")
		}
		content, _ = os.ReadFile(bobFile)
		if !strings.Contains(string(content), "blazingly fast") {
			t.Errorf("Bob's file should contain Rust joke")
		}

		// Verify Charlie -> main.md (Default)
		charlieFile := filepath.Join(tempDir, "solutions", "day02", "Charlie", "main.md")
		if _, err := os.Stat(charlieFile); os.IsNotExist(err) {
			t.Errorf("Charlie should have main.md")
		}
	})
}
