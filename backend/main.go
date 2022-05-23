package main

import (
	"os"
	"log"
	"database/sql"

	"loganlieou.xyz/omelon_backend/utils"

	_ "github.com/go-sql-driver/mysql"
)

/*
Enviornment Variables
DB_HOST
is of the form
user:pass@tcp(ip:port)/dbname
*/

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", os.Getenv("DB_HOST"))

	if err != nil {
		log.Fatal(err)
	}
	
	// CURD
	utils.PrintWarehouse(db, "Austin")
}

