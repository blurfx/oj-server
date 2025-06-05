package model

import "time"

type Attachment struct {
	ID          uint      `db:"id"`
	Name        string    `db:"name"`
	ContentSize uint      `db:"content_size"`
	URL         string    `db:"url"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	DeletedAt   time.Time `db:"deleted_at"`
}
