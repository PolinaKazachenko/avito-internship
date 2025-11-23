package users

import (
	"context"

	"avito-internship/internal/app/dto"
	"avito-internship/internal/app/repository/postgresql"
	"avito-internship/internal/app/repository/tx_facade"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TxFacade struct {
	pool *pgxpool.Pool
}

func NewTxFacade(pool *pgxpool.Pool) *TxFacade {
	return &TxFacade{pool: pool}
}

func (f *TxFacade) AddUsers(ctx context.Context, users []*dto.User) (int64, error) {
	facade := tx_facade.NewTxFacade(f.pool)
	var result int64

	err := facade.Execute(ctx, func(tx pgx.Tx) error {
		userRepo := postgresql.NewUserRepository(tx)
		var err error

		result, err = userRepo.AddUsers(ctx, users)
		return err
	})
	return result, err
}

func (f *TxFacade) UserListByTeamName(ctx context.Context, teamName string) ([]*dto.User, error) {
	facade := tx_facade.NewTxFacade(f.pool)
	var result []*dto.User

	err := facade.Execute(ctx, func(tx pgx.Tx) error {
		userRepo := postgresql.NewUserRepository(tx)
		var err error

		result, err = userRepo.UserListByTeamName(ctx, teamName)
		return err
	})
	return result, err
}

func (f *TxFacade) UserSetActive(ctx context.Context, userID string, active bool) (*dto.User, error) {
	facade := tx_facade.NewTxFacade(f.pool)
	var result *dto.User

	err := facade.Execute(ctx, func(tx pgx.Tx) error {
		userRepo := postgresql.NewUserRepository(tx)
		var err error

		result, err = userRepo.UserSetActive(ctx, userID, active)
		return err
	})
	return result, err
}

func (f *TxFacade) GetTeamUserListByUserID(ctx context.Context, authorID string) ([]*dto.User, error) {
	facade := tx_facade.NewTxFacade(f.pool)
	var result []*dto.User

	err := facade.Execute(ctx, func(tx pgx.Tx) error {
		userRepo := postgresql.NewUserRepository(tx)
		var err error

		result, err = userRepo.GetTeamUserListByUserID(ctx, authorID)
		return err
	})
	return result, err
}
