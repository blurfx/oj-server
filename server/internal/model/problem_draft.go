package model

import "time"

type ProblemDraft struct {
	ID             uint      `db:"id"`
	Title          string    `db:"title"`
	Description    string    `db:"description"`
	Spoiler        string    `db:"spoiler"`
	TimeLimit      int64     `db:"time_limit"`
	MemoryLimit    int64     `db:"memory_limit"`
	DraftMappingID uint      `db:"draft_mapping_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	DeletedAt      time.Time `db:"deleted_at"`
}
