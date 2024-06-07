package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

func main() {
    // Database connection string
    dsn := "root:example@tcp(db:3306)/testdb"

    // Open database connection
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Error opening database: %v\n", err)
    }

    // Test the database connection
    err = db.Ping()
    if err != nil {
        log.Fatalf("Error pinging database: %v\n", err)
    }

    fmt.Println("Successfully connected to the MySQL database!")

    // Close the database connection
    defer db.Close()
}
