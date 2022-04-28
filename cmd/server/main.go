package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4/middleware"
	"github.com/orderoutofchaos/oj-server/internal/dao"
	"github.com/orderoutofchaos/oj-server/internal/datasource/mysql"
	handler2 "github.com/orderoutofchaos/oj-server/internal/handler"
	"os"
)

func main() {
	if err := godotenv.Load(".env.local", ".env"); err != nil {
		panic(err)
	}

	readerConfig := mysql.Config{
		User:   os.Getenv("READER_DB_USER"),
		Passwd: os.Getenv("READER_DB_PASS"),
		Host:   os.Getenv("READER_DB_HOST"),
		Port:   os.Getenv("READER_DB_PORT"),
		DBName: os.Getenv("READER_DB_NAME"),
	}
	reader := mysql.NewMySQL(readerConfig)
	writerConfig := mysql.Config{
		User:   os.Getenv("WRITER_DB_USER"),
		Passwd: os.Getenv("WRITER_DB_PASS"),
		Host:   os.Getenv("WRITER_DB_HOST"),
		Port:   os.Getenv("WRITER_DB_PORT"),
		DBName: os.Getenv("WRITER_DB_NAME"),
	}
	writer := mysql.NewMySQL(writerConfig)
	dao.InitRepo(reader, writer)

	e := echo.New()
	secret := ""
	e.Use(middleware.Logger())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(secret))))
	e.Validator = &handler2.RequestValidator{
		Validator: validator.New(),
	}
	handler2.InitV1Handler(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}
