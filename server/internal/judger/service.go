package judger

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/blurfx/fxoj/internal/dao"
	"github.com/blurfx/fxoj/internal/model"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			err := s.processNextSubmission(ctx)
			if err != nil {
				log.Printf("Error processing submission: %v", err)
			}
			time.Sleep(time.Second)
		}
	}
}

func (s *Service) processNextSubmission(ctx context.Context) error {
	repo := dao.GetRepo()
	tx := repo.Writer().MustBegin()
	defer tx.Rollback()

	submission := model.Submission{}
	err := tx.Get(
		&submission,
		`
		SELECT *
		FROM submissions
		WHERE status = $1
		ORDER BY id ASC
		FOR UPDATE SKIP LOCKED
		LIMIT 1
	`, model.SubmissionStatusPending)

	fmt.Println(submission, err)
	if err == sql.ErrNoRows {
		return nil
	}

	if err != nil {
		return fmt.Errorf("get submission: %w", err)
	}

	tx.MustExec(`
		UPDATE submissions
		SET status = $1, updated_at = NOW()
		WHERE id = $2
	`, model.SubmissionStatusRunning, submission.ID)

	go s.judgeSubmission(ctx, submission)

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func (s *Service) judgeSubmission(_ context.Context, submission model.Submission) {
	// TODO: Implement actual judging logic
	log.Printf("Judging submission %d", submission.ID)
}
