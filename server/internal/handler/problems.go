package handler

import (
	"database/sql"
	"net/http"

	"github.com/blurfx/fxoj/internal/dao"
	"github.com/blurfx/fxoj/internal/model"
	"github.com/labstack/echo/v4"
)

type V1GetProblemsRequest struct {
	Size uint `json:"size"`
	Page uint `json:"page"`
}

type V1GetProblemsResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

func V1GetProblems(c echo.Context, req *V1GetProblemsRequest) Response {
	repo := dao.GetRepo()
	if req.Size == 0 {
		req.Size = 10
	}
	if req.Page == 0 {
		req.Page = 1
	}
	problems := []model.Problem{}
	err := repo.Reader().Select(
		&problems,
		"SELECT id, title FROM problems LIMIT $1 OFFSET $2",
		req.Size,
		(req.Page-1)*req.Size,
	)
	if err != nil {
		panic(err)
	}

	data := make([]V1GetProblemsResponse, len(problems))
	for i, problem := range problems {
		data[i] = V1GetProblemsResponse{
			ID:    problem.ID,
			Title: problem.Title,
		}
	}

	return Response{
		Code: http.StatusOK,
		Data: data,
	}
}

type V1GetProblemRequest struct {
	ID uint `json:"id"`
}

type V1GetProblemResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Spoiler     string `json:"spoiler"`
	TimeLimit   uint   `json:"time_limit"`
	MemoryLimit uint   `json:"memory_limit"`
}

func V1GetProblem(c echo.Context, _ *struct{}) Response {
	repo := dao.GetRepo()
	id := c.Param("id")
	problem := model.Problem{}
	err := repo.Reader().Get(&problem, "SELECT id, title, description, spoiler, time_limit, memory_limit FROM problems WHERE id = $1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return Response{
				Code: http.StatusNotFound,
			}
		}
		panic(err)
	}

	return Response{
		Code: http.StatusOK,
		Data: V1GetProblemResponse{
			ID:          problem.ID,
			Title:       problem.Title,
			Description: problem.Description,
			Spoiler:     problem.Spoiler,
			TimeLimit:   problem.TimeLimit,
			MemoryLimit: problem.MemoryLimit,
		},
	}
}
