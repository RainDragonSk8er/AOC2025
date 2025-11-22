package aoc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type JSONResponse struct {
	Event    string            `json:"event"`
	OwnerID  int               `json:"owner_id"`
	Day1TS   int64             `json:"day1_ts"`
	NumDays  int               `json:"num_days"`
	Members  map[string]Member `json:"members"`
}

type Member struct {
	ID                 int                    `json:"id"`
	Name               *string                `json:"name"` // pointer because name can be null
	Stars              int                    `json:"stars"`
	LastStarTS         int64                  `json:"last_star_ts"`
	LocalScore         int                    `json:"local_score"`
	CompletionDayLevel map[string]interface{} `json:"completion_day_level"`
}

type Client struct {
	SessionCookie string
	LeaderboardID string
	UserAgent     string
}

func NewClient(cookie, leaderboardID string) *Client {
	return &Client{
		SessionCookie: cookie,
		LeaderboardID: leaderboardID,
		UserAgent:     "github.com/RainDragonSk8er/AOC2025 by aoc-bot@example.com",
	}
}

func (c *Client) FetchLeaderboard() (*JSONResponse, error) {
	apiUrl := fmt.Sprintf("https://adventofcode.com/2025/leaderboard/private/view/%s.json", c.LeaderboardID)
	cookieHeader := fmt.Sprintf("session=%s", c.SessionCookie)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", cookieHeader)
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var leaderboard JSONResponse
	if err := json.Unmarshal(body, &leaderboard); err != nil {
		return nil, err
	}

	return &leaderboard, nil
}
