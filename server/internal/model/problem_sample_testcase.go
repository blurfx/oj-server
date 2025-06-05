package model

import "time"

type ProblemSampleTestcase struct {
	ID             uint      `db:"id"`
	ProblemID      uint      `db:"problem_id"`
	ProblemDraftID uint      `db:"draft_id"`
	Input          string    `db:"input"`
	Output         string    `db:"output"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	DeletedAt      time.Time `db:"deleted_at"`
}
