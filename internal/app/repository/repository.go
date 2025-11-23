package repository

import (
	"context"

	"avito-internship/internal/app/dto"
)

type (
	TeamRepository interface {
		AddTeam(ctx context.Context, team *dto.Team) (bool, error)
	}

	UserRepository interface {
		AddUsers(ctx context.Context, users []*dto.User) (int64, error)
		UserListByTeamName(ctx context.Context, teamName string) ([]*dto.User, error)
		UserSetActive(ctx context.Context, userID string, active bool) (*dto.User, error)
		GetTeamUserListByUserID(ctx context.Context, authorID string) ([]*dto.User, error)
	}

	PullRequestRepository interface {
		AddPullRequest(ctx context.Context, pullRequest *dto.PullRequest) (bool, error)
		MergePullRequest(ctx context.Context, pullRequest *dto.PullRequest) (bool, error)
		GetPullRequest(ctx context.Context, pullRequestID string) (*dto.PullRequest, error)
		ReassignReviewer(ctx context.Context, pullRequestID string, oldReviewer string, newReviewer string) (*dto.PullRequest, error)
		PullRequestListByReviewerID(ctx context.Context, reviewerID string) ([]*dto.PullRequest, error)
	}
)
