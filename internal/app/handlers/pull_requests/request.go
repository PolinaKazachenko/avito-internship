package pull_requests

import "avito-internship/internal/app/dto"

type CreateRequest struct {
	PullRequestID   string `json:"pull_request_id" binding:"required,min=1"`
	PullRequestName string `json:"pull_request_name" binding:"required,min=1"`
	AuthorID        string `json:"author_id" binding:"required,min=1"`
}

type MergeRequest struct {
	PullRequestID string `json:"pull_request_id" binding:"required,min=1"`
}

type ReassignRequest struct {
	PullRequestID string `json:"pull_request_id" binding:"required,min=1"`
	OldReviewerID string `json:"old_reviewer_id" binding:"required,min=1"`
}

func convertCreateRequest(req *CreateRequest) *dto.PullRequest {
	return &dto.PullRequest{
		ID:       req.PullRequestID,
		Name:     req.PullRequestName,
		AuthorID: req.AuthorID,
	}
}
