package pull_requests

import (
	"context"

	"avito-internship/internal/app/dto"
	"avito-internship/internal/app/repository/postgresql"
	"avito-internship/internal/app/repository/tx_facade"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type txFacade struct {
	pool *pgxpool.Pool
}

func NewTxFacade(pool *pgxpool.Pool) *txFacade {
	return &txFacade{pool: pool}
}

func (f *txFacade) AddPullRequest(ctx context.Context, pullRequest *dto.PullRequest) (bool, error) {
	facade := tx_facade.NewTxFacade(f.pool)
	var result bool

	err := facade.Execute(ctx, func(tx pgx.Tx) error {
		prRepo := postgresql.NewPullRequestRepository(tx)
		var err error
		result, err = prRepo.AddPullRequest(ctx, pullRequest)
		return err
	})
	return result, err
}

func (f *txFacade) MergePullRequest(ctx context.Context, pullRequest *dto.PullRequest) (bool, error) {
	facade := tx_facade.NewTxFacade(f.pool)
	var result bool

	err := facade.Execute(ctx, func(tx pgx.Tx) error {
		prRepo := postgresql.NewPullRequestRepository(tx)
		var err error
		result, err = prRepo.MergePullRequest(ctx, pullRequest)
		return err
	})
	return result, err
}

func (f *txFacade) GetPullRequest(ctx context.Context, pullRequestID string) (*dto.PullRequest, error) {
	var result *dto.PullRequest
	facade := tx_facade.NewTxFacade(f.pool)

	err := facade.Execute(ctx, func(tx pgx.Tx) error {
		prRepo := postgresql.NewPullRequestRepository(tx)
		var err error
		result, err = prRepo.GetPullRequest(ctx, pullRequestID)
		return err
	})
	return result, err
}

func (f *txFacade) ReassignReviewer(ctx context.Context, pullRequestID string, oldReviewer string, newReviewer string) (*dto.PullRequest, error) {
	facade := tx_facade.NewTxFacade(f.pool)
	var result *dto.PullRequest

	err := facade.Execute(ctx, func(tx pgx.Tx) error {
		prRepo := postgresql.NewPullRequestRepository(tx)
		var err error
		result, err = prRepo.ReassignReviewer(ctx, pullRequestID, oldReviewer, newReviewer)
		return err
	})
	return result, err
}

func (f *txFacade) PullRequestListByReviewerID(ctx context.Context, reviewerID string) ([]*dto.PullRequest, error) {

	facade := tx_facade.NewTxFacade(f.pool)
	var result []*dto.PullRequest

	err := facade.Execute(ctx, func(tx pgx.Tx) error {
		prRepo := postgresql.NewPullRequestRepository(tx)
		var err error
		result, err = prRepo.PullRequestListByReviewerID(ctx, reviewerID)
		return err
	})
	return result, err
}
