package model

type ProblemDraftMapping struct {
	ID        uint `db:"id"`
	ProblemID uint `db:"problem_id"`
	DraftID   uint `db:"draft_id"`
}
