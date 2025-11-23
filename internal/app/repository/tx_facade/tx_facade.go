package tx_facade

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type txFacade struct {
	pool *pgxpool.Pool
}

func NewTxFacade(pool *pgxpool.Pool) *txFacade {
	return &txFacade{pool: pool}
}

func (f *txFacade) Execute(ctx context.Context, task func(tx pgx.Tx) error) (err error) {
	tx, err := f.pool.Begin(ctx)
	defer tx.Rollback(ctx) // nolint:errcheck
	if err = task(tx); err != nil {
		return
	}
	return tx.Commit(ctx)
}
