package pull_requests

import (
	"time"

	"avito-internship/internal/app/dto"
)

type PullRequest struct {
	ID          string   `json:"pull_request_id"`
	Name        string   `json:"pull_request_name"`
	AuthorID    string   `json:"author_id"`
	Status      string   `json:"status"`
	ReviewerIDs []string `json:"assigned_reviewers"`
}

type FullPullRequest struct {
	*PullRequest
	MergedAt time.Time `json:"merged_at,omitempty"`
}

func convertPullRequest(pr *dto.PullRequest) *PullRequest {
	return &PullRequest{
		ID:          pr.ID,
		Name:        pr.Name,
		AuthorID:    pr.AuthorID,
		Status:      convertPulRequestStatus(pr.Status),
		ReviewerIDs: pr.ReviewerIDs,
	}
}

func convertFullPullRequest(pr *dto.PullRequest) *FullPullRequest {
	return &FullPullRequest{
		PullRequest: convertPullRequest(pr),
		MergedAt:    pr.MergedAt,
	}
}

func convertPulRequestStatus(status dto.PullRequestStatus) string {
	switch status {
	case dto.Open:
		return "OPEN"
	case dto.Merged:
		return "MERGED"
	default:
		return "UNKNOWN"
	}
}
