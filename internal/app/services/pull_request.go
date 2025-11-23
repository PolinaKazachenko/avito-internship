package services

import (
	"context"
	"errors"
	"slices"
	"time"

	"avito-internship/internal/app/dto"
	"avito-internship/internal/app/repository"
	"avito-internship/internal/logger"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

const (
	numOfAssignReviewers = 2
)

type PullRequestService struct {
	pullRequestRepo repository.PullRequestRepository
	userRepo        repository.UserRepository
}

func NewPullRequestService(
	pullRequestRepo repository.PullRequestRepository,
	userRepo repository.UserRepository,
) *PullRequestService {
	return &PullRequestService{
		pullRequestRepo: pullRequestRepo,
		userRepo:        userRepo,
	}
}

func (f *PullRequestService) AddPullRequest(ctx context.Context, pullRequest *dto.PullRequest) (*dto.PullRequest, error) {
	teamUsers, err := f.userRepo.GetTeamUserListByUserID(ctx, pullRequest.AuthorID)
	if err != nil {
		logger.ErrorKV(ctx, "AddPullRequest.GetTeamUserListByUserID", "error", err.Error())
		return nil, ErrDatabase
	}
	if len(teamUsers) == 0 {
		return nil, ErrNotFound
	}
	pullRequest.ReviewerIDs = getAssignedUserIDs(teamUsers, pullRequest.AuthorID)
	pullRequest.Status = dto.Open
	if _, err = f.pullRequestRepo.AddPullRequest(ctx, pullRequest); err != nil {
		logger.ErrorKV(ctx, "AddPullRequest.AddPullRequest", "error", err.Error())
		var pgxErr *pgconn.PgError
		if errors.As(err, &pgxErr) && pgxErr.Code == UniqueViolationCode {
			return nil, ErrAlreadyExists
		}
		return nil, ErrDatabase
	}
	return pullRequest, nil
}

func (f *PullRequestService) MergePullRequest(ctx context.Context, pullRequestID string) (*dto.PullRequest, error) {
	pullRequest, err := f.pullRequestRepo.GetPullRequest(ctx, pullRequestID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		logger.ErrorKV(ctx, "MergePullRequest.GetPullRequest", "error", err.Error())
		return nil, ErrDatabase
	}
	if pullRequest.Status == dto.Merged {
		return pullRequest, nil
	}
	pullRequest.Status = dto.Merged
	pullRequest.MergedAt = time.Now()
	_, err = f.pullRequestRepo.MergePullRequest(ctx, pullRequest)
	if err != nil {
		logger.ErrorKV(ctx, "MergePullRequest.MergePullRequest", "error", err.Error())
		return nil, ErrDatabase
	}
	return pullRequest, nil
}

func (f *PullRequestService) ReassignReviewer(ctx context.Context, pullRequestID string, oldReviewerID string) (*dto.PullRequest, string, error) {
	pullRequest, err := f.pullRequestRepo.GetPullRequest(ctx, pullRequestID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, "", ErrNotFound
		}
		logger.ErrorKV(ctx, "MergePullRequest.GetPullRequest", "error", err.Error())
		return nil, "", ErrDatabase
	}
	if pullRequest.Status == dto.Merged {
		return nil, "", ErrPrAlreadyMerged
	}
	oldReviewerIdx := slices.Index(pullRequest.ReviewerIDs, oldReviewerID)
	if oldReviewerIdx == -1 {
		return nil, "", ErrReviewerNotAssignedToPR
	}
	teamUsers, err := f.userRepo.GetTeamUserListByUserID(ctx, pullRequest.AuthorID)
	if err != nil {
		logger.ErrorKV(ctx, "MergePullRequest.GetTeamUserListByUserID", "error", err.Error())
		return nil, "", ErrDatabase
	}
	if len(teamUsers) == 0 {
		return nil, "", ErrNotFound
	}
	newReviewerID, err := getNewReviewerID(oldReviewerID, pullRequest, teamUsers)
	if err != nil {
		return nil, "", err
	}
	result, err := f.pullRequestRepo.ReassignReviewer(ctx, pullRequestID, oldReviewerID, newReviewerID)
	if err != nil {
		logger.ErrorKV(ctx, "MergePullRequest.ReassignReviewer", "error", err.Error())
		return nil, "", ErrDatabase
	}
	return result, newReviewerID, nil
}

func getNewReviewerID(
	oldReviewerID string,
	pullRequest *dto.PullRequest,
	teamUsers []*dto.User,
) (string, error) {
	slices.Sort(pullRequest.ReviewerIDs)
	for _, teamUser := range teamUsers {
		if teamUser.ID == pullRequest.AuthorID || teamUser.ID == oldReviewerID || !teamUser.IsActive {
			continue
		}
		_, ok := slices.BinarySearch(pullRequest.ReviewerIDs, teamUser.ID)
		if !ok {
			return teamUser.ID, nil
		}
	}
	return "", ErrNoActiveCandidates
}

func getAssignedUserIDs(teamUsers []*dto.User, authorID string) []string {
	result := make([]string, 0, numOfAssignReviewers)
	idx := 0
	for _, user := range teamUsers {
		if idx == numOfAssignReviewers {
			break
		}
		if user.ID == authorID || !user.IsActive {
			continue
		}
		result = append(result, user.ID)
		idx++
	}
	return result
}
