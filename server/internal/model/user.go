package model

import "time"

type UserRole string

const (
	UserRoleAdmin  UserRole = "admin"
	UserRoleAuthor UserRole = "author"
	UserRoleUser   UserRole = "user"
)

type User struct {
	ID          uint      `db:"id"`
	Username    string    `db:"username"`
	Password    string    `db:"password"`
	Role        UserRole  `db:"role"`
	LastLoginAt time.Time `db:"last_login_at"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	DeletedAt   time.Time `db:"deleted_at"`
}
