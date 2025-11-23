package postgresql

import (
	"context"

	"avito-internship/internal/app/dto"
	"avito-internship/internal/app/repository"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
)

type UserRepository struct {
	tx pgx.Tx
}

// NewUserRepository ...
func NewUserRepository(tx pgx.Tx) repository.UserRepository {
	return &UserRepository{tx: tx}
}

// AddUsers ...
func (r *UserRepository) AddUsers(ctx context.Context, users []*dto.User) (int64, error) {
	cnt, err := r.tx.CopyFrom(
		ctx,
		pgx.Identifier{"users"},
		[]string{"id", "username", "team_name"},
		pgx.CopyFromSlice(len(users), func(i int) ([]interface{}, error) {
			return []interface{}{users[i].ID, users[i].UserName, users[i].TeamName}, nil
		}),
	)
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

// UserListByTeamName ...
func (r *UserRepository) UserListByTeamName(ctx context.Context, teamName string) ([]*dto.User, error) {
	rawResults := make([]User, 0)
	if err := pgxscan.Select(ctx, r.tx, &rawResults, `SELECT * FROM users WHERE team_name = $1`, teamName); err != nil {
		return nil, err
	}
	results := make([]*dto.User, 0, len(rawResults))
	for _, result := range rawResults {
		results = append(results, result.toDomainObject())
	}
	return results, nil
}

// UserSetActive ...
func (r *UserRepository) UserSetActive(ctx context.Context, userID string, active bool) (*dto.User, error) {
	q := `UPDATE users SET is_active = $1, updated_at = now() WHERE id = $2 
		  RETURNING id, username, team_name, is_active, created_at, updated_at`
	user := &User{}

	if err := pgxscan.Get(ctx, r.tx, user, q, active, userID); err != nil {
		return nil, err
	}
	return user.toDomainObject(), nil
}

func (r *UserRepository) GetTeamUserListByUserID(ctx context.Context, userID string) ([]*dto.User, error) {
	q := `SELECT u2.id, u2.username, u2.team_name, u2.is_active, u2.created_at, u2.updated_at FROM users as u1
		  LEFT JOIN users as u2 ON u1.team_name = u2.team_name
		  WHERE u1.id = $1;`
	rawUsers := make([]*User, 0)

	if err := pgxscan.Select(ctx, r.tx, &rawUsers, q, userID); err != nil {
		return nil, err
	}
	results := make([]*dto.User, 0, len(rawUsers))
	for _, result := range rawUsers {
		results = append(results, result.toDomainObject())
	}
	return results, nil
}
