package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
    "time"

    _ "github.com/go-sql-driver/mysql"
    "github.com/ajtroup1/speakeasy/types"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

var db *sql.DB

func main() {
    err := godotenv.Load("./config/vars.env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    connectToDb()

    router := gin.Default()

    router.GET("/user/:email", getUserByEmail)

    router.Run("localhost:8080")
}

func connectToDb() {
    host := os.Getenv("HOST")
    user := os.Getenv("USER")
    password := os.Getenv("PASSWORD")
    dbname := os.Getenv("DBNAME")
    port := os.Getenv("PORT")

    dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

    var err error
    db, err = sql.Open("mysql", dbURI)
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatalf("Error pinging the database: %v", err)
    }

    log.Println("Connected to the database")
}

func getUserByEmail(c *gin.Context) {
    email := c.Param("email")

    // Simulate fetching user data based on email
    

    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
