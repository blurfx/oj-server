package mysql

import (
    "database/sql"
    "fmt"
    "github.com/go-sql-driver/mysql"
    "time"
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
    config := &mysql.Config{
        User:                 c.User,
        Passwd:               c.Passwd,
        Net:                  "tcp",
        Addr:                 fmt.Sprintf("%s:%s", c.Host, c.Port),
        DBName:               c.DBName,
        Collation:            "utf8mb4_unicode_ci",
        ParseTime:            true,
        AllowNativePasswords: true,
    }
    return config.FormatDSN()
}

func NewMySQL(config Config) *sql.DB {
    db, err := sql.Open("mysql", config.FormatDSN())
    if err != nil {
        panic(err)
    }
    return db
}
