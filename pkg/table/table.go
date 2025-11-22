package table

import (
	"fmt"
	"sort"
	"strings"

	"github.com/RainDragonSk8er/AOC2025/pkg/aoc"
)

// Generate creates a markdown ASCII table from the leaderboard data.
func Generate(l *aoc.JSONResponse) string {
	var sb strings.Builder

	sb.WriteString("| Rank | Name | Stars | Local Score |\n")
	sb.WriteString("| :--- | :--- | :---: | :---: |\n")

	// Convert map to slice for sorting
	type entry struct {
		Member aoc.Member
	}
	var entries []entry
	for _, m := range l.Members {
		entries = append(entries, entry{Member: m})
	}

	// Sort by LocalScore (descending), then Stars (descending), then Name (ascending)
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Member.LocalScore != entries[j].Member.LocalScore {
			return entries[i].Member.LocalScore > entries[j].Member.LocalScore
		}
		if entries[i].Member.Stars != entries[j].Member.Stars {
			return entries[i].Member.Stars > entries[j].Member.Stars
		}
		nameI := "Anonymous"
		if entries[i].Member.Name != nil {
			nameI = *entries[i].Member.Name
		}
		nameJ := "Anonymous"
		if entries[j].Member.Name != nil {
			nameJ = *entries[j].Member.Name
		}
		return nameI < nameJ
	})

	for i, e := range entries {
		name := "Anonymous"
		if e.Member.Name != nil {
			name = *e.Member.Name
		}
		// Escape pipes in names just in case
		name = strings.ReplaceAll(name, "|", "\\|")
		
		sb.WriteString(fmt.Sprintf("| %d | %s | %d | %d |\n", i+1, name, e.Member.Stars, e.Member.LocalScore))
	}

	return sb.String()
}
