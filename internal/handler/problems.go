package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/orderoutofchaos/oj-server/internal/dao"
)

type V1GetProblemsRequest struct {
	Size uint `json:"size"`
	Page uint `json:"page"`
}

type ProblemListItem struct {
	ID    uint   `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
}

func V1GetProblems(req *V1GetProblemsRequest, c echo.Context) Response {
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
