package services

import (
	"context"
	"errors"

	"avito-internship/internal/app/dto"
	"avito-internship/internal/app/repository"
	"avito-internship/internal/logger"
	"github.com/jackc/pgx/v4"
)

type UserService struct {
	userRepository        repository.UserRepository
	pullRequestRepository repository.PullRequestRepository
}

func NewUserService(
	userRepository repository.UserRepository,
	pullRequestRepository repository.PullRequestRepository,
) *UserService {
	return &UserService{
		userRepository:        userRepository,
		pullRequestRepository: pullRequestRepository,
	}
}

func (s *UserService) SetIsActive(ctx context.Context, userID string, isActive bool) (*dto.User, error) {
	user, err := s.userRepository.UserSetActive(ctx, userID, isActive)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		logger.ErrorKV(ctx, "UserService.SetIsActive", "error", err)
		return nil, ErrDatabase
	}
	return user, nil
}

func (s *UserService) GetReview(ctx context.Context, userID string) (*dto.UserPullRequests, error) {
	pullRequests, err := s.pullRequestRepository.PullRequestListByReviewerID(ctx, userID)
	if err != nil {
		logger.ErrorKV(ctx, "UserService.GetReview", "error", err)
		return nil, ErrDatabase
	}
	return &dto.UserPullRequests{
		UserID:       userID,
		PullRequests: pullRequests,
	}, nil
}
