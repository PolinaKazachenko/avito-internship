package postgresql

import (
	"time"

	"avito-internship/internal/app/dto"
	"github.com/lib/pq"
)

const (
	openStatus    = "OPEN"
	mergedStatus  = "MERGED"
	unknownStatus = "UNKNOWN"
)

type User struct {
	ID        string      `db:"id"`
	UserName  string      `db:"username"`
	TeamName  string      `db:"team_name"`
	IsActive  bool        `db:"is_active"`
	CreatedAt time.Time   `db:"created_at"`
	UpdatedAt pq.NullTime `db:"updated_at"`
}

type PullRequest struct {
	ID          string         `db:"id"`
	Name        string         `db:"name"`
	AuthorID    string         `db:"author_id"`
	Status      string         `db:"status"`
	ReviewerIDs pq.StringArray `db:"reviewer_ids"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   pq.NullTime    `db:"updated_at"`
	MergedAt    pq.NullTime    `db:"merged_at"`
}

func (u User) toDomainObject() *dto.User {
	var updatedAt time.Time
	if u.UpdatedAt.Valid {
		updatedAt = u.UpdatedAt.Time
	}
	return &dto.User{
		ID:        u.ID,
		UserName:  u.UserName,
		TeamName:  u.TeamName,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt,
		UpdatedAt: updatedAt,
	}
}

func (p PullRequest) toDomainObject() *dto.PullRequest {
	var updatedAt time.Time
	var mergedAt time.Time
	if p.UpdatedAt.Valid {
		updatedAt = p.UpdatedAt.Time
	}
	if p.MergedAt.Valid {
		mergedAt = p.MergedAt.Time
	}
	return &dto.PullRequest{
		ID:          p.ID,
		Name:        p.Name,
		AuthorID:    p.AuthorID,
		Status:      pullRequestStatusToDomain(p.Status),
		ReviewerIDs: p.ReviewerIDs,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   updatedAt,
		MergedAt:    mergedAt,
	}
}

func pullRequestStatusToDomain(status string) dto.PullRequestStatus {
	switch status {
	case openStatus:
		return dto.Open
	case mergedStatus:
		return dto.Merged
	default:
		return dto.Unknown
	}
}

func pullRequestStatusFromDomain(status dto.PullRequestStatus) string {
	switch status {
	case dto.Open:
		return openStatus
	case dto.Merged:
		return mergedStatus
	default:
		return unknownStatus
	}
}
