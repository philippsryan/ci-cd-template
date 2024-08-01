package api

import (
	"database/sql"
	"github.com/labstack/echo/v4"
)

type GetDatabase func() (*sql.DB, error)

type TodoContext struct {
	echo.Context
	GetDatabase
}
