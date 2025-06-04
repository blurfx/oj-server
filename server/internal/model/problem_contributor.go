package model

type ProblemContributorType string

const (
	ProblemContributorTypeReviewer ProblemContributorType = "reviewer"
)

type ProblemContributor struct {
	ID        uint                   `db:"id"`
	UserID    uint                   `db:"user_id"`
	ProblemID uint                   `db:"problem_id"`
	Type      ProblemContributorType `db:"type"`
}
