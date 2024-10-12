package database

import (
    "database/sql"
    "log"

    "github.com/sijms/go-ora/v2"

)

var DB *sql.DB
var conString string = go_ora.BuildUrl("localhost",1521,"orcl","gowtham","1234",nil)
func InitDB() {
    var err error
    // Replace with your Oracle DB connection string
    DB, err = sql.Open("oracle", conString)
    if err != nil {
        log.Fatalf("Error opening Oracle connection: %v", err)
    }

    // Test the connection
    if err = DB.Ping(); err != nil {
        log.Fatalf("Unable to connect to Oracle DB: %v", err)
    }
    log.Println("Connected to Oracle Database!")
}