package postgresql

import (
	"context"

	"avito-internship/internal/app/dto"
	"avito-internship/internal/app/repository"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
)

type PullRequestRepository struct {
	tx pgx.Tx
}

// NewPullRequestRepository ...
func NewPullRequestRepository(tx pgx.Tx) repository.PullRequestRepository {
	return &PullRequestRepository{tx: tx}
}

// AddPullRequest ...
func (r *PullRequestRepository) AddPullRequest(ctx context.Context, pullRequest *dto.PullRequest) (bool, error) {
	q := `INSERT INTO pull_requests(id, name, author_id, status, reviewer_ids) VALUES($1, $2, $3, $4, $5)`
	result, err := r.tx.Exec(ctx, q,
		pullRequest.ID,
		pullRequest.Name,
		pullRequest.AuthorID,
		pullRequestStatusFromDomain(pullRequest.Status),
		pq.StringArray(pullRequest.ReviewerIDs),
	)
	if err != nil {
		return false, err
	}
	return result.RowsAffected() > 0, nil
}

// MergePullRequest ...
func (r *PullRequestRepository) MergePullRequest(ctx context.Context, pullRequest *dto.PullRequest) (bool, error) {
	q := `UPDATE pull_requests SET status = $1, merged_at = $2 WHERE id = $3 AND merged_at IS NULL`
	result, err := r.tx.Exec(ctx, q,
		pullRequestStatusFromDomain(pullRequest.Status),
		pullRequest.MergedAt,
		pullRequest.ID)
	if err != nil {
		return false, err
	}
	return result.RowsAffected() > 0, err
}

// PullRequestListByReviewerID ...
func (r *PullRequestRepository) PullRequestListByReviewerID(ctx context.Context, reviewerID string) ([]*dto.PullRequest, error) {
	rawResults := make([]PullRequest, 0)
	q := `SELECT id, name, author_id, status, reviewer_ids, created_at, updated_at FROM pull_requests 
		  WHERE pull_requests.reviewer_ids && ARRAY[$1]`

	if err := pgxscan.Select(ctx, r.tx, &rawResults, q, reviewerID); err != nil {
		return nil, err
	}
	results := make([]*dto.PullRequest, 0, len(rawResults))
	for _, result := range rawResults {
		results = append(results, result.toDomainObject())
	}
	return results, nil
}

// ReassignReviewer ...
func (r *PullRequestRepository) ReassignReviewer(ctx context.Context, pullRequestID string, oldReviewer string, newReviewer string) (*dto.PullRequest, error) {
	var result *PullRequest
	q := `UPDATE pull_requests SET reviewer_ids = array_replace(reviewer_ids, $1, $2) WHERE id = $3
          RETURNING id, name, author_id, status, reviewer_ids, created_at, updated_at, merged_at`
	if err := pgxscan.Get(ctx, r.tx, result, q, oldReviewer, newReviewer, pullRequestID); err != nil {
		return nil, err
	}
	return result.toDomainObject(), nil
}

// GetPullRequest ...
func (r *PullRequestRepository) GetPullRequest(ctx context.Context, pullRequestID string) (*dto.PullRequest, error) {
	q := `SELECT id, name, author_id, status, reviewer_ids, created_at, updated_at, merged_at 
			 FROM pull_requests WHERE id = $1`
	rawResult := &PullRequest{}

	err := pgxscan.Get(ctx, r.tx, rawResult, q, pullRequestID)
	if err != nil {
		return nil, err
	}
	return rawResult.toDomainObject(), err
}
