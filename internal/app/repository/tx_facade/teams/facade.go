package teams

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

func (f *TxFacade) AddTeam(ctx context.Context, team *dto.Team) (bool, error) {
	facade := tx_facade.NewTxFacade(f.pool)
	var result bool
	err := facade.Execute(ctx, func(tx pgx.Tx) error {
		teamRepo := postgresql.NewTeamRepository(tx)
		userRepo := postgresql.NewUserRepository(tx)
		var err error
		result, err = teamRepo.AddTeam(ctx, team)
		if err != nil {
			return err
		}
		_, err = userRepo.AddUsers(ctx, team.Members)
		return err
	})
	return result, err
}
