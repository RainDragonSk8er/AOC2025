package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/RainDragonSk8er/AOC2025/pkg/aoc"
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

	// Create day directory
	dayDir := fmt.Sprintf("solutions/day%02d", targetDay)
	if err := os.MkdirAll(dayDir, 0755); err != nil {
		log.Fatal("Error creating day directory:", err)
	}

	// Iterate members
	for _, member := range leaderboard.Members {
		name := "Anonymous"
		if member.Name != nil {
			name = *member.Name
		}

		folderName := name
		if mapped, ok := mapping[name]; ok {
			folderName = mapped
		}
		// Sanitize folder name
		folderName = strings.ReplaceAll(folderName, " ", "_")
		folderName = strings.ReplaceAll(folderName, "/", "_")

		memberDir := filepath.Join(dayDir, folderName)
		if err := os.MkdirAll(memberDir, 0755); err != nil {
			log.Printf("Error creating directory for %s: %v", name, err)
			continue
		}

		// Smart Language Detection
		extensions := detectPreviousLanguage(targetDay, folderName)
		if len(extensions) == 0 {
			extensions = []string{".md"} // Default
		}

		for _, ext := range extensions {
			filename := "main" + ext
			filePath := filepath.Join(memberDir, filename)
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				content := ""
				if ext == ".md" {
					content = fmt.Sprintf("# Solution for Day %d\n\nTODO: Write solution", targetDay)
				} else {
					content = fmt.Sprintf("// Solution for Day %d", targetDay)
				}
				os.WriteFile(filePath, []byte(content), 0644)
				fmt.Printf("Created %s\n", filePath)
			}
		}
	}
}

func detectPreviousLanguage(currentDay int, folderName string) []string {
	if currentDay <= 1 {
		return nil
	}
	prevDayDir := fmt.Sprintf("solutions/day%02d/%s", currentDay-1, folderName)
	files, err := os.ReadDir(prevDayDir)
	if err != nil {
		return nil
	}

	knownExts := map[string]bool{
		".py": true, ".go": true, ".rs": true, ".js": true, ".ts": true,
		".cpp": true, ".c": true, ".java": true, ".rb": true, ".lua": true,
		".jl": true,
	}

	var found []string
	for _, file := range files {
		if !file.IsDir() {
			ext := filepath.Ext(file.Name())
			if knownExts[ext] {
				found = append(found, ext)
			}
		}
	}
	return found
}
