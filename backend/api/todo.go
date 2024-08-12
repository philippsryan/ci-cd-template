package api

import (
	"net/http"
	"todoapp/db"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	Id         int    `json:"id"`
	Belongs_To string `json:"belongs_to"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Done       bool   `json:"done"`
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
		result, err := database.Query("SELECT Id, Title, Body FROM Todo WHERE BelongsTo = ?", user)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.String(http.StatusInternalServerError, "")
		}

		todos := make([]Todo, 0)

		for result.Next() {
			todo := Todo{Title: "", Belongs_To: user, Body: "", Id: -1}
			result.Scan(&todo.Id, &todo.Title, &todo.Body)
			todos = append(todos, todo)
		}

		return c.JSON(http.StatusOK, todos)

	} else {
		c.Logger().Error(c.Request().Form)
		return c.String(http.StatusBadRequest, user+" does not exist!")
	}
}

func GetTodoById(c echo.Context) error {
	id_to_get := c.Param("id")

	database, err := db.CreateDatabaseConnection()

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get database connection")
	}

	rows, err := database.Query("SELECT Id, Title, Body, Done FROM Todo WHERE Id = ?", id_to_get)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to query database")
	}

	todo := Todo{Title: "", Done: false, Body: "", Id: -1}

	if rows.Next() {
		rows.Scan(&todo.Id, &todo.Title, &todo.Body, &todo.Done)
	} else {
		return c.String(http.StatusNotFound, "Could not find todo with that id")
	}

	return c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c echo.Context) error {
	id_to_update := c.Param("id")
	updated_todo := Todo{}
	err := (&echo.DefaultBinder{}).BindBody(c, &updated_todo)

	if err != nil {
		return c.String(http.StatusBadRequest, "Failed to parse body as a todo")
	}

	database, err := db.CreateDatabaseConnection()

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get database connection")
	}

	result, err := database.Exec("UPDATE Todo SET Title = ?, Body = ?, Done = ? WHERE Id = ?", updated_todo.Title, updated_todo.Body, updated_todo.Done, id_to_update)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get database connection")
	}

	rows, err := result.RowsAffected()

	if rows > 0 {
		return c.JSON(http.StatusOK, updated_todo)
	} else {
		return c.String(http.StatusNotFound, "Failed to find todo with that id")
	}
}
