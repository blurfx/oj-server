package postgres

import (
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	User    string
	Passwd  string
	Host    string
	Port    string
	DBName  string
	MaxConn int
	Timeout time.Duration
}

func (c *Config) FormatDSN() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", c.User, c.Passwd, c.Host, c.Port, c.DBName)
}

func NewPostgres(config Config) *sqlx.DB {
	db, err := sqlx.Open("pgx", config.FormatDSN())
	if err != nil {
		panic(err)
	}
	return db
}
