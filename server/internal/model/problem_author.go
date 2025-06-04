package model

type ProblemAuthor struct {
	ID        uint `db:"id"`
	UserID    uint `db:"user_id"`
	ProblemID uint `db:"problem_id"`
}
