package dto

import "time"

type PullRequestStatus = int

const (
	Unknown PullRequestStatus = iota
	Open
	Merged
)

type PullRequest struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	AuthorID    string            `json:"author_id"`
	Status      PullRequestStatus `json:"status"`
	ReviewerIDs []string          `json:"reviewers"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	MergedAt    time.Time         `json:"merged_at"`
}

func ConvertStatusToString(pullRequestStatus PullRequestStatus) string {
	switch pullRequestStatus {
	case Open:
		return "OPEN"
	case Merged:
		return "MERGED"
	default:
		return "UNKNOWN"
	}
}
