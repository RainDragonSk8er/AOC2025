package table

import (
	"encoding/json"
	"strings"
	"testing"

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

	output := Generate(&leaderboard)

	// Expected rank order: Anonymous (120), Alice (100), Bob (80)
	expectedLines := []string{
		"| Rank | Name | Stars | Local Score |",
		"| :--- | :--- | :---: | :---: |",
		"| 1 | Anonymous | 12 | 120 |",
		"| 2 | Alice | 10 | 100 |",
		"| 3 | Bob | 8 | 80 |",
	}

	for _, line := range expectedLines {
		if !strings.Contains(output, line) {
			t.Errorf("Expected output to contain:\n%s\nBut got:\n%s", line, output)
		}
	}
}
