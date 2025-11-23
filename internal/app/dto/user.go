package dto

import "time"

type User struct {
	ID        string    `json:"id"`
	UserName  string    `json:"user_name"`
	IsActive  bool      `json:"is_active"`
	TeamName  string    `json:"team_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserPullRequests struct {
	UserID       string         `json:"user_id"`
	PullRequests []*PullRequest `json:"pull_requests"`
}
