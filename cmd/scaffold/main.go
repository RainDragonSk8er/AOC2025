package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"

	"github.com/RainDragonSk8er/AOC2025/pkg/aoc"
	"github.com/RainDragonSk8er/AOC2025/pkg/scaffold"
	"github.com/joho/godotenv"
)

func main() {
	dayPtr := flag.Int("day", 0, "Day number to scaffold (required)")
	flag.Parse()

	targetDay := *dayPtr
	if targetDay == 0 {
		now := time.Now()
		// Logic for AOC 2025 (Shortened: Dec 1 - Dec 12)
		// We also allow Nov 30 to scaffold Day 1 early if desired, but user asked for strict midnight appearance?
		// User comment: "Please just scaffold 30. November -> 11 or 12 th of December"
		// This implies:
		// Nov 30 -> Day 1? Or just strict dates?
		// "just scaffold 30. November -> 11 or 12 th of December"
		// I will interpret this as:
		// If today is Nov 30, scaffold Day 1.
		// If today is Dec 1-12, scaffold that day.

		_, month, day := now.Date()
		if month == time.November && day == 30 {
			targetDay = 1
		} else if month == time.December && day >= 1 && day <= 12 {
			targetDay = day
		} else {
			// Not in the active window, do nothing silently (so cron doesn't spam errors)
			return
		}
	}

	// Load .env
	godotenv.Load() // Ignore error

	apiCookie := os.Getenv("API_COOKIE")
	apiLeaderboard := os.Getenv("API_LEADERBOARD")
	if apiCookie == "" || apiLeaderboard == "" {
		log.Fatal("API_COOKIE and API_LEADERBOARD must be set")
	}

	// Load mapping
	mapping := make(map[string]string)
	mappingData, err := os.ReadFile("config/mapping.json")
	if err == nil {
		json.Unmarshal(mappingData, &mapping)
	} else {
		log.Println("Warning: config/mapping.json not found or readable")
	}

	// Fetch leaderboard
	client := aoc.NewClient(apiCookie, apiLeaderboard)
	leaderboard, err := client.FetchLeaderboard()
	if err != nil {
		log.Fatal("Error fetching leaderboard:", err)
	}

	// Run scaffolding
	// We use "." as rootDir since we are running from the project root
	if err := scaffold.Run(".", targetDay, leaderboard, mapping); err != nil {
		log.Fatal(err)
	}
}
