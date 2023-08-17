package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/databasename?parseTime=true")
    if err != nil {
        log.Fatal(err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    { // Query a single user
        var (
			index	  		string
            id        		string
            username  		string
            password  		string
			name      		string
			email     		string
			department		string
			designation 	string
			region			string
			role			string
			date_created	string
        )

        query := "SELECT * FROM users WHERE username=?"
        if err := db.QueryRow(query, "bane").Scan(&index, &id, &username, &password, &name, &email, &department, &designation, &region, &role, &date_created); err != nil {
            log.Fatal(err)
        }

        fmt.Println(id, username, password, date_created)
    }
}