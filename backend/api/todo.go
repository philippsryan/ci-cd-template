package api

import (
	"net/http"
	"todoapp/db"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	Belongs_To string `json:"belongs_to"`
	Title      string `json:"title"`
	Body       string `json:"body"`
}

func CreateTodo(c echo.Context) error {
	var todo Todo
	err := c.Bind(&todo)
	if err != nil {
		c.String(http.StatusBadRequest, "Could not parse todo to create")
	}

	database, _ := db.CreateDatabaseConnection()

	if userAlreadyExists(todo.Belongs_To, database, c) {
		_, err := database.Exec("INSERT INTO Todo (BelongsTo, Title, Body) VALUES (?, ?, ?)", todo.Belongs_To, todo.Title, todo.Body)

		if err != nil {
			c.Logger().Error(err.Error())
			return c.String(http.StatusInternalServerError, "")
		}

		return c.String(http.StatusOK, "Todo created")
	} else {
		return c.String(http.StatusBadRequest, todo.Belongs_To+" does not exist!")
	}
}

func GetTodo(c echo.Context) error {
	user := c.QueryParam("user")

	database, _ := db.CreateDatabaseConnection()

	if userAlreadyExists(user, database, c) {
		result, err := database.Query("SELECT Title, Body FROM Todo WHERE BelongsTo = ?", user)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.String(http.StatusInternalServerError, "")
		}

		todos := make([]Todo, 0)

		for result.Next() {
			todo := Todo{Title: "", Belongs_To: user, Body: ""}
			result.Scan(&todo.Title, &todo.Body)
			todos = append(todos, todo)
		}

		return c.JSON(http.StatusOK, todos)

	} else {
		c.Logger().Error(c.Request().Form)
		return c.String(http.StatusBadRequest, user+" does not exist!")
	}
}
