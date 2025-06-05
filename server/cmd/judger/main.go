package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/blurfx/fxoj/internal/dao"
	"github.com/blurfx/fxoj/internal/datasource/postgres"
	"github.com/blurfx/fxoj/internal/judger"
	"github.com/blurfx/fxoj/internal/utils"
)

func main() {
	utils.LoadEnv()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	writerConfig := postgres.Config{
		User:   os.Getenv("WRITER_DB_USER"),
		Passwd: os.Getenv("WRITER_DB_PASS"),
		Host:   os.Getenv("WRITER_DB_HOST"),
		Port:   os.Getenv("WRITER_DB_PORT"),
		DBName: os.Getenv("WRITER_DB_NAME"),
	}
	writer := postgres.NewPostgres(writerConfig)
	dao.InitRepo(writer, writer)

	judgeService := judger.NewService()

	go func() {
		if err := judgeService.Start(ctx); err != nil {
			log.Printf("Judge service error: %v", err)
			cancel()
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	log.Println("Shutting down judge server...")
}
