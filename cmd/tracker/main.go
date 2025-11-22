package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/RainDragonSk8er/AOC2025/pkg/aoc"
	"github.com/RainDragonSk8er/AOC2025/pkg/table"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		// It's okay if .env doesn't exist (e.g. in Docker), we might use actual env vars
		log.Println("Warning: Error loading .env file, checking environment variables directly.")
	}

	// Retrieve environment variables
	apiCookie := os.Getenv("API_COOKIE")
	apiLeaderboard := os.Getenv("API_LEADERBOARD")

	if apiCookie == "" || apiLeaderboard == "" {
		log.Fatal("API_COOKIE and API_LEADERBOARD must be set")
	}

	client := aoc.NewClient(apiCookie, apiLeaderboard)
	leaderboard, err := client.FetchLeaderboard()
	if err != nil {
		log.Fatal("Error fetching leaderboard:", err)
	}

	// Load Theme
	theme := table.Theme{
		Bar: table.BarConfig{Filled: "#", Empty: "."},
		Emoticons: []table.EmoticonConfig{
			{Threshold: 0, Icon: "(╯°□°)╯︵ ┻━┻"},
		},
	}
	themeData, err := os.ReadFile("config/theme.json")
	if err == nil {
		json.Unmarshal(themeData, &theme)
	} else {
		log.Println("Warning: config/theme.json not found, using default theme")
	}

	// Generate ASCII table
	asciiTable := table.Generate(leaderboard, theme)

	// Update README.md
	readmePath := "README.md"
	content, err := os.ReadFile(readmePath)
	if err != nil {
		log.Fatal("Error reading README.md:", err)
	}

	// We want to replace everything after a certain marker, or just append/replace the table section.
	// For simplicity, let's assume we want to replace the entire "Leaderboard" section if it exists,
	// or just append it.
	// A better approach for this specific repo structure is to look for a marker.
	// Let's define a marker in the README.
	marker := "<!-- LEADERBOARD_START -->"
	endMarker := "<!-- LEADERBOARD_END -->"

	text := string(content)
	startIdx := -1
	endIdx := -1

	// Simple string search
	for i := 0; i < len(text)-len(marker); i++ {
		if text[i:i+len(marker)] == marker {
			startIdx = i
			break
		}
	}

	if startIdx != -1 {
		// Find end marker
		for i := startIdx; i < len(text)-len(endMarker); i++ {
			if text[i:i+len(endMarker)] == endMarker {
				endIdx = i + len(endMarker)
				break
			}
		}
	}

	newContent := ""
	tableSection := fmt.Sprintf("%s\n\n%s\n%s", marker, asciiTable, endMarker)

	if startIdx != -1 && endIdx != -1 {
		// Replace existing section
		newContent = text[:startIdx] + tableSection + text[endIdx:]
	} else {
		// Append to end if not found
		newContent = text + "\n\n## Leaderboard\n" + tableSection
	}

	if newContent != text {
		err = os.WriteFile(readmePath, []byte(newContent), 0644)
		if err != nil {
			log.Fatal("Error writing README.md:", err)
		}
		fmt.Println("README.md updated.")
	} else {
		fmt.Println("No changes to leaderboard.")
	}
}
