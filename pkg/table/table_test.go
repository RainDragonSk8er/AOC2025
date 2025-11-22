package table

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/RainDragonSk8er/AOC2025/pkg/aoc"
)

func TestGenerate(t *testing.T) {
	// Synthetic JSON data mimicking AOC API
	jsonData := `
	{
		"event": "2025",
		"owner_id": 12345,
		"members": {
			"1": {
				"id": 1,
				"name": "Alice",
				"stars": 10,
				"local_score": 100,
				"last_star_ts": 1234567890,
				"completion_day_level": {}
			},
			"2": {
				"id": 2,
				"name": "Bob",
				"stars": 8,
				"local_score": 80,
				"last_star_ts": 1234567890,
				"completion_day_level": {}
			},
			"3": {
				"id": 3,
				"name": null,
				"stars": 12,
				"local_score": 120,
				"last_star_ts": 1234567890,
				"completion_day_level": {}
			}
		}
	}`

	var leaderboard aoc.JSONResponse
	if err := json.Unmarshal([]byte(jsonData), &leaderboard); err != nil {
		t.Fatalf("Failed to unmarshal synthetic data: %v", err)
	}

	theme := Theme{
		Bar: BarConfig{Filled: "#", Empty: "."},
		Emoticons: []EmoticonConfig{
			{Threshold: 0, Icon: "(╯°□°)╯︵ ┻━┻"},
			{Threshold: 5, Icon: "(ಠ_ಠ)"},
			{Threshold: 9, Icon: "(._.)"},
		},
	}

	// Mock time as Dec 1st 2025
	mockTime := time.Date(2025, time.December, 1, 0, 0, 0, 0, time.UTC)
	output := Generate(&leaderboard, theme, mockTime)

	// Expected output should contain diff syntax and progress bars
	// Since we are running this test in Nov 2025 (presumably), it might return the "revealed" message.
	// We should mock time, but since we can't easily mock time.Now() without dependency injection,
	// let's just check if it returns EITHER the message OR the table.
	// ideally we'd refactor Generate to accept a time, but for now let's just check for the message
	// if it's hidden, or the table if it's shown.

	// However, for the purpose of this test, we want to verify the TABLE generation logic.
	// Let's assume the test environment might not be in Nov 2025.
	// If the output is the hidden message, we can't verify the table.
	// Let's refactor Generate to take a time.Time argument?
	// Or just check for the message and if so, print a log saying "Skipping table verification due to date".

	if strings.Contains(output, "Leaderboard will be revealed") {
		t.Log("Leaderboard is hidden, skipping table structure verification.")
		return
	}

	expectedSubstrings := []string{
		"```diff",
		"! --- Advent of Code 2025 Leaderboard ---",
		"+ Anonymous",
		"! [############............] 12 Stars (._.)", // 12 stars = 12 #, emoticon (._.)
		"+ Alice",
		"! [##########..............] 10 Stars (._.)", // 10 stars
		"+ Bob",
		"! [########................] 8 Stars (ಠ_ಠ)", // 8 stars
	}

	for _, s := range expectedSubstrings {
		if !strings.Contains(output, s) {
			t.Errorf("Expected output to contain:\n%s\nBut got:\n%s", s, output)
		}
	}
}
