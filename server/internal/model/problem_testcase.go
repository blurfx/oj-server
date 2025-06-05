package model

import "time"

type ProblemTestcase struct {
	ID                 uint      `db:"id"`
	ProblemID          uint      `db:"problem_id"`
	ProblemDraftID     uint      `db:"draft_id"`
	InputAttachmentID  string    `db:"input_attachment_id"`
	OutputAttachmentID string    `db:"output_attachment_id"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
	DeletedAt          time.Time `db:"deleted_at"`
}
