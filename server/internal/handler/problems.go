package handler

import (
	"net/http"

	"github.com/blurfx/fxoj/internal/dao"
	"github.com/labstack/echo/v4"
)

type V1GetProblemsRequest struct {
	Size uint `json:"size"`
	Page uint `json:"page"`
}

type ProblemListItem struct {
	ID    uint   `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
}

func V1GetProblems(c echo.Context, req *V1GetProblemsRequest) Response {
	repo := dao.GetRepo()
	if req.Size == 0 {
		req.Size = 10
	}
	if req.Page == 0 {
		req.Page = 1
	}
	rows, err := repo.Reader().Query(
		"SELECT id, title FROM problem LIMIT $1 OFFSET $2",
		req.Size,
		(req.Page-1)*req.Size,
	)
	if err != nil {
		panic(err)
	}

	problems := make([]ProblemListItem, 0)
	for rows.Next() {
		var problem ProblemListItem
		rows.ScanStruct(&problem)
		problems = append(problems, problem)
	}
	return Response{
		Code: http.StatusOK,
		Data: problems,
	}
}

type V1GetProblemRequest struct {
	ID uint `json:"id"`
}

type Problem struct {
	ID          uint   `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Spoiler     string `db:"spoiler" json:"spoiler"`
	TimeLimit   uint   `db:"time_limit" json:"time_limit"`
	MemoryLimit uint   `db:"memory_limit" json:"memory_limit"`
}

func V1GetProblem(c echo.Context, _ *struct{}) Response {
	repo := dao.GetRepo()
	id := c.Param("id")
	rows, err := repo.Reader().Query("SELECT id, title, description, spoiler, time_limit, memory_limit FROM problem WHERE id = $1", id)

	if err != nil {
		panic(err)
	}

	problems := make([]Problem, 0)
	for rows.Next() {
		var problem Problem
		rows.ScanStruct(&problem)
		problems = append(problems, problem)
	}

	if len(problems) == 0 {
		return Response{
			Code: http.StatusNotFound,
		}
	}

	return Response{
		Code: http.StatusOK,
		Data: problems[0],
	}
}
