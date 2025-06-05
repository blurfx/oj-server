package main

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/blurfx/fxoj/internal/dao"
	"github.com/blurfx/fxoj/internal/datasource/postgres"
	"github.com/blurfx/fxoj/internal/handler"
)

func loadEnv() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	if env != "test" {
		godotenv.Load(".env.local") //nolint
	}
	godotenv.Load(".env." + env) //nolint
	godotenv.Load()              //nolint
}

func main() {
	loadEnv()

	readerConfig := postgres.Config{
		User:   os.Getenv("READER_DB_USER"),
		Passwd: os.Getenv("READER_DB_PASS"),
		Host:   os.Getenv("READER_DB_HOST"),
		Port:   os.Getenv("READER_DB_PORT"),
		DBName: os.Getenv("READER_DB_NAME"),
	}
	reader := postgres.NewPostgres(readerConfig)
	writerConfig := postgres.Config{
		User:   os.Getenv("WRITER_DB_USER"),
		Passwd: os.Getenv("WRITER_DB_PASS"),
		Host:   os.Getenv("WRITER_DB_HOST"),
		Port:   os.Getenv("WRITER_DB_PORT"),
		DBName: os.Getenv("WRITER_DB_NAME"),
	}
	writer := postgres.NewPostgres(writerConfig)
	dao.InitRepo(reader, writer)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))))
	e.Validator = &handler.RequestValidator{
		Validator: validator.New(),
	}
	handler.InitV1Handler(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}
