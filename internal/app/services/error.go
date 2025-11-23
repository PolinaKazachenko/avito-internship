package services

import "errors"

const (
	UniqueViolationCode = "23505"
)

var (
	ErrDatabase = errors.New("database error")

	ErrNotFound = errors.New("not found error")

	ErrAlreadyExists = errors.New("already exists error")

	ErrPrAlreadyMerged = errors.New("cannot reassign on merged PR")

	ErrReviewerNotAssignedToPR = errors.New("reviewer is not assigned to this PR")

	ErrNoActiveCandidates = errors.New("no active replacement candidate in team")
)
