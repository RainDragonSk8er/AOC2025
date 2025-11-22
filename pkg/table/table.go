package table

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/RainDragonSk8er/AOC2025/pkg/aoc"
)

type Theme struct {
	Bar       BarConfig        `json:"bar"`
	Emoticons []EmoticonConfig `json:"emoticons"`
}

type BarConfig struct {
	Filled string `json:"filled"`
	Empty  string `json:"empty"`
}

type EmoticonConfig struct {
	Threshold int    `json:"threshold"`
	Icon      string `json:"icon"`
}

// Generate creates a markdown ASCII table from the leaderboard data.
func Generate(l *aoc.JSONResponse, theme Theme, now time.Time) string {
	// Date Check: Hide before Dec 1st (unless it's 2024/2025 transition testing, but strict rule requested)
	// Actually, user said "not be present in the README.md before the first of december".
	if now.Month() != time.December || now.Day() < 1 {
		// If not December (and not past it, e.g. Jan), hide it.
		// But we want it to persist in Jan 2026.
		// So: If Year < 2025, hide. If Year == 2025 and Month < Dec, hide.
		// Simple check: If it's Nov 2025, return empty.
		if now.Year() == 2025 && now.Month() < time.December {
			return "Leaderboard will be revealed on December 1st!"
		}
	}

	var sb strings.Builder

	sb.WriteString("```diff\n")
	sb.WriteString("! --- Advent of Code 2025 Leaderboard ---\n\n")

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

	maxStars := 24 // 12 days * 2 stars
	barWidth := 24 // 1 char per star

	for _, e := range entries {
		name := "Anonymous"
		if e.Member.Name != nil {
			name = *e.Member.Name
		}

		stars := e.Member.Stars
		if stars > maxStars {
			stars = maxStars
		}

		// Progress Bar
		filled := stars
		empty := barWidth - filled
		bar := strings.Repeat(theme.Bar.Filled, filled) + strings.Repeat(theme.Bar.Empty, empty)

		// Emoticon
		emoticon := ""
		// Find the highest threshold met
		currentThreshold := -1
		for _, emo := range theme.Emoticons {
			if stars >= emo.Threshold && emo.Threshold > currentThreshold {
				emoticon = emo.Icon
				currentThreshold = emo.Threshold
			}
		}
		if emoticon == "" && len(theme.Emoticons) > 0 {
			emoticon = theme.Emoticons[0].Icon // Fallback

		}

		sb.WriteString(fmt.Sprintf("+ %s\n", name))
		sb.WriteString(fmt.Sprintf("! [%s] %d Stars %s\n\n", bar, e.Member.Stars, emoticon))
	}
	sb.WriteString("```\n")

	return sb.String()
}
