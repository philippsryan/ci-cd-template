package api

import (
	"database/sql"
	"github.com/labstack/echo/v4"
)

type TodoContext struct {
	echo.Context
	*sql.DB
}
