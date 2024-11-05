package main

import (
	"net/http"
	"os"
	"strconv"
	"todoapp/api"
	"todoapp/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	args := os.Args

	command := args[1]

	println(command, "== migrations:", command == "migrations")

	if command == "migrations" {

		database, _ := db.CreateDatabaseConnection()
		completed_migrations, err := db.RunMigrations(database)

		if err != nil {
			panic("Completed: " + strconv.FormatInt(int64(completed_migrations), 10) + " but ran into error: " + err.Error())
		}

		println("Migrations are done! Completed", completed_migrations)
		return
	}

	if command == "server" {
		e := echo.New()
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}))
		e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				todo_context := &api.TodoContext{Context: c, GetDatabase: db.CreateDatabaseConnection}
				return next(todo_context)
			}
		})

		e.POST("/user", api.CreateUser)
		e.GET("/user", api.GetAllUsers)

		e.GET("/todos", api.GetTodo)
		e.POST("/todos", api.CreateTodo)
		e.GET("/todos/:id", api.GetTodoById)

		e.Logger.Fatal(e.Start(":8000"))
	} else {
		println("please provide a command to run")
	}

}
