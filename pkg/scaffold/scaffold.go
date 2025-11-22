package scaffold

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/RainDragonSk8er/AOC2025/pkg/aoc"
)

// Run executes the scaffolding logic for a specific day.
// rootDir is the base directory where "solutions/" will be created.
func Run(rootDir string, day int, leaderboard *aoc.JSONResponse, mapping map[string]string) error {
	// Create day directory
	dayDir := filepath.Join(rootDir, "solutions", fmt.Sprintf("day%02d", day))
	if err := os.MkdirAll(dayDir, 0755); err != nil {
		return fmt.Errorf("error creating day directory: %w", err)
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
		extensions := detectPreviousLanguage(rootDir, day, folderName)
		if len(extensions) == 0 {
			extensions = []string{".md"} // Default
		}

		for _, ext := range extensions {
			filename := "main" + ext
			filePath := filepath.Join(memberDir, filename)
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				content := GetTemplate(ext, day)
				if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
					log.Printf("Error writing file %s: %v", filePath, err)
				} else {
					fmt.Printf("Created %s\n", filePath)
				}
			}
		}
	}
	return nil
}

func detectPreviousLanguage(rootDir string, currentDay int, folderName string) []string {
	if currentDay <= 1 {
		return nil
	}
	prevDayDir := filepath.Join(rootDir, "solutions", fmt.Sprintf("day%02d", currentDay-1), folderName)
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
