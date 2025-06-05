package handler

import (
	"database/sql"
	"encoding/hex"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/scrypt"

	"github.com/blurfx/fxoj/internal/dao"
	"github.com/blurfx/fxoj/internal/model"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const salt = ""

func encodeHash(value string) string {
	key, _ := scrypt.Key([]byte(value), []byte(salt), 32768, 8, 1, 32)
	return hex.EncodeToString(key)
}

func V1Login(c echo.Context, req *LoginRequest) Response {
	repo := dao.GetRepo()
	var user model.User
	err := repo.Reader().Get(
		&user,
		"SELECT id, username FROM users WHERE username = $1 AND password = $2",
		req.Username, encodeHash(req.Password),
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return Response{
				Code:  http.StatusUnauthorized,
				Error: ErrInvalidCredential,
			}
		}
		panic(err)
	}

	sess, err := session.Get("session", c)
	if err != nil {
		return Response{
			Code:  http.StatusInternalServerError,
			Error: ErrSession,
		}
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["user_id"] = user.ID
	sess.Values["username"] = user.Username
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		log.Printf("auth login: %v", err)
		return Response{
			Code:  http.StatusInternalServerError,
			Error: ErrInternal,
		}
	}

	return Response{
		Code: http.StatusOK,
	}
}

func V1Logout(c echo.Context, _ *struct{}) Response {
	sess, err := session.Get("session", c)
	if err != nil {
		return Response{
			Code:  http.StatusInternalServerError,
			Error: ErrSession,
		}
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		log.Printf("auth logout: %v", err)
		return Response{
			Code:  http.StatusInternalServerError,
			Error: ErrInternal,
		}
	}
	return Response{
		Code: http.StatusOK,
	}
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func V1Register(c echo.Context, req *RegisterRequest) Response {
	repo := dao.GetRepo()
	var user model.User
	err := repo.Reader().Get(
		&user,
		"SELECT id FROM users WHERE username = $1",
		req.Username,
	)
	if err == nil {
		return Response{
			Code:  http.StatusBadRequest,
			Error: ErrUserAlreadyExists,
		}
	}
	if err != sql.ErrNoRows {
		panic(err)
	}

	repo.Writer().MustExec(
		"INSERT INTO users (username, password) VALUES ($1, $2)",
		req.Username,
		encodeHash(req.Password),
	)
	return Response{
		Code: http.StatusOK,
	}
}
