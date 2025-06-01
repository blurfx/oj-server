package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/blurfx/fxoj/pkg/sq"
	_ "github.com/jackc/pgx/v5/stdlib"
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

func NewPostgres(config Config) *sq.DB {
	db, err := sql.Open("pgx", config.FormatDSN())
	if err != nil {
		panic(err)
	}
	return sq.NewDb(db)
}
