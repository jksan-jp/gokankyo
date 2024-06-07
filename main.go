package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// DB connection string
const dsn = "root:example@tcp(db:3306)/testdb"

// User represents a user in the database
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// Connect to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}
	defer db.Close()

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v\n", err)
	}
	fmt.Println("Successfully connected to the MySQL database!")

	// Create a new Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/sample", func(c echo.Context) error {
		var users []User

		// Query the database
		rows, err := db.Query("SELECT id, name FROM users")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Error querying database",
			})
		}
		defer rows.Close()

		// Iterate over the rows
		for rows.Next() {
			var user User
			if err := rows.Scan(&user.ID, &user.Name); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": "Error scanning row",
				})
			}
			users = append(users, user)
		}

		// Return the users as JSON
		return c.JSON(http.StatusOK, users)
	})

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
