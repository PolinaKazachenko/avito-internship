package services

import (
	"context"
	"errors"

	"avito-internship/internal/app/dto"
	"avito-internship/internal/app/repository"
	"avito-internship/internal/logger"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type TeamService struct {
	teamRepository repository.TeamRepository
	userRepository repository.UserRepository
}

func NewTeamService(
	teamRepository repository.TeamRepository,
	userRepository repository.UserRepository,
) *TeamService {
	return &TeamService{
		teamRepository: teamRepository,
		userRepository: userRepository,
	}
}

func (s *TeamService) Add(ctx context.Context, team *dto.Team) (*dto.Team, error) {
	_, err := s.teamRepository.AddTeam(ctx, team)
	if err != nil {
		logger.ErrorKV(ctx, "TeamService.Add", "error", err)
		var pgxErr *pgconn.PgError
		if errors.As(err, &pgxErr) && pgxErr.Code == UniqueViolationCode {
			return nil, ErrAlreadyExists
		}
		return nil, ErrDatabase
	}
	return team, nil
}

func (s *TeamService) Get(ctx context.Context, teamName string) (*dto.Team, error) {
	members, err := s.userRepository.UserListByTeamName(ctx, teamName)
	if err != nil {
		logger.ErrorKV(ctx, "TeamService.Get", "error", err)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, ErrDatabase
	}
	if len(members) == 0 {
		return nil, ErrNotFound
	}
	return &dto.Team{
		Name:    teamName,
		Members: members,
	}, nil
}
