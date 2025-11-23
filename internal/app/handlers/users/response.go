package users

import "avito-internship/internal/app/dto"

type User struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	TeamName string `json:"team_name"`
	IsActive bool   `json:"is_active"`
}

type SetIsActiveResponse struct {
	User *User `json:"user"`
}

type PullRequest struct {
	ID       string `json:"pull_request_id"`
	Name     string `json:"pull_request_name"`
	AuthorID string `json:"author_id"`
	Status   string `json:"status"`
}

type GetReviewResponse struct {
	UserID       string         `json:"user_id"`
	PullRequests []*PullRequest `json:"pull_requests"`
}

func convertIsActiveResponse(user *dto.User) *SetIsActiveResponse {
	if user == nil {
		return nil
	}
	return &SetIsActiveResponse{
		User: &User{
			UserID:   user.ID,
			UserName: user.UserName,
			TeamName: user.TeamName,
			IsActive: user.IsActive,
		},
	}
}

func convertGetReviewResponse(user *dto.UserPullRequests) *GetReviewResponse {
	if user == nil {
		return nil
	}
	pullRequests := make([]*PullRequest, 0, len(user.PullRequests))
	for _, pr := range user.PullRequests {
		pullRequests = append(pullRequests, &PullRequest{
			ID:       pr.ID,
			Name:     pr.Name,
			AuthorID: pr.AuthorID,
			Status:   dto.ConvertStatusToString(pr.Status),
		})
	}
	return &GetReviewResponse{
		PullRequests: pullRequests,
		UserID:       user.UserID,
	}
}
