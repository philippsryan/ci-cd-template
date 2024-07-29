package main

import (
	"os"
	"strconv"
	"todoapp/api"
	"todoapp/db"

	"github.com/labstack/echo/v4"
)

func main() {

	args := os.Args

	command := args[1]

	println(command, "== migrations:", command == "migrations")

	if command == "migrations" {

		database := db.CreateDatabaseConnection()
		completed_migrations, err := db.RunMigrations(database)

		if err != nil {
			panic("Completed: " + strconv.FormatInt(int64(completed_migrations), 10) + " but ran into error: " + err.Error())
		}

		println("Migrations are done! Completed", completed_migrations)
		return
	}

	if command == "server" {
		e := echo.New()

		database := db.CreateDatabaseConnection()
		e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				todo_context := &api.TodoContext{Context: c, DB: database}
				return next(todo_context)
			}
		})

		e.POST("/user", api.CreateUser)
		e.GET("/user", api.GetAllUsers)

		e.GET("/todos", api.GetTodo)
		e.POST("/todos", api.CreateTodo)

		e.Logger.Fatal(e.Start(":8000"))
	} else {
		println("please provide a command to run")
	}

}
