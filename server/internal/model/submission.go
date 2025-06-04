package model

import "time"

type SubmissionStatus string

const (
	SubmissionStatusPending             SubmissionStatus = "pending"
	SubmissionStatusRunning             SubmissionStatus = "running"
	SubmissionStatusAccepted            SubmissionStatus = "accepted"
	SubmissionStatusWrongAnswer         SubmissionStatus = "wrong_answer"
	SubmissionStatusTimeLimitExceeded   SubmissionStatus = "time_limit_exceeded"
	SubmissionStatusMemoryLimitExceeded SubmissionStatus = "memory_limit_exceeded"
	SubmissionStatusRuntimeError        SubmissionStatus = "runtime_error"
	SubmissionStatusCompileError        SubmissionStatus = "compile_error"
	SubmissionStatusSystemError         SubmissionStatus = "system_error"
)

type SubmissionVisibility string

const (
	SubmissionVisibilityPublic     SubmissionVisibility = "public"
	SubmissionVisibilityPrivate    SubmissionVisibility = "private"
	SubmissionVisibilitySolvedOnly SubmissionVisibility = "solved_only"
)

type Submission struct {
	ID          uint                 `db:"id"`
	Code        string               `db:"code"`
	TimeLimit   int64                `db:"time_limit"`
	MemoryLimit int64                `db:"memory_limit"`
	Status      SubmissionStatus     `db:"status"`
	Visibility  SubmissionVisibility `db:"visibility"`
	LanguageID  uint                 `db:"language_id"`
	ProblemID   uint                 `db:"problem_id"`
	UserID      uint                 `db:"user_id"`
	CreatedAt   time.Time            `db:"created_at"`
	UpdatedAt   time.Time            `db:"updated_at"`
	DeletedAt   time.Time            `db:"deleted_at"`
}
