package postgresql

import (
	"context"

	"avito-internship/internal/app/dto"
	"github.com/jackc/pgx/v4"
)

type TeamRepository struct {
	tx pgx.Tx
}

func NewTeamRepository(tx pgx.Tx) *TeamRepository {
	return &TeamRepository{tx: tx}
}

func (r *TeamRepository) AddTeam(ctx context.Context, team *dto.Team) (bool, error) {
	result, err := r.tx.Exec(ctx, `INSERT INTO teams(name) VALUES($1)`, team.Name)
	if err != nil {
		return false, err
	}
	return result.RowsAffected() > 0, nil
}
