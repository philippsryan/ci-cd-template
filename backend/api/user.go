package api

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func userAlreadyExists(name string, db *sql.DB, c echo.Context) bool {
	result := db.QueryRow("SELECT Name FROM User WHERE Name = ?;", name)
	if result.Err() != nil {
		c.Logger().Error(result.Err().Error())
	}

	var name_in_db string

	err := result.Scan(&name_in_db)
	if err != nil {
		return false
	}

	return name_in_db == name
}

func CreateUser(c echo.Context) error {
	todo_context := c.(*TodoContext)
	database, err := todo_context.GetDatabase()
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to load data")
	}

	user_name := c.FormValue("username")

	if userAlreadyExists(user_name, database, c) {
		return c.String(http.StatusBadRequest, user_name+" already exists silly")
	}

	insert_row, err := database.Query("INSERT INTO User VALUES (?);", user_name)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to insert user into database")
	}

	insert_row.Close()

	database.Close()

	return c.String(http.StatusOK, "Added "+user_name)
}

func GetAllUsers(c echo.Context) error {
	todo_context := c.(*TodoContext)
	database, err := todo_context.GetDatabase()
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to load data")
	}

	user_rows, err := database.Query("SELECT Name FROM User;")

	if err != nil {
		return c.String(http.StatusInternalServerError, "Could not get users")
	}

	users := make([]string, 0)

	for user_rows.Next() {
		name := ""

		user_rows.Scan(&name)
		users = append(users, name)
	}

	database.Close()
	return c.JSON(http.StatusOK, users)
}
