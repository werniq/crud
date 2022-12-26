package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func (app *application) createConnection() (*sql.DB, error) {
	var err error
	var port1 int
	
	host := getVarFromDotenv("host")
	port := getVarFromDotenv("port")
	user := getVarFromDotenv("user")
	db_name := getVarFromDotenv("db_name")
	// Creating dsn for database connection
	port1, err = strconv.Atoi(port)
	
	if err != nil {
		app.errorLog.Println(err)
		return nil, err
	}
	
	app.cfg.dbDsn = fmt.Sprintf("host=%s port=%d user=%s db_name=%s sslmode=disabled", host, port1, user, db_name)
	app.cfg.port, err = strconv.Atoi(port)

	if err != nil {
		app.errorLog.Println(err)
		return nil, err
	}


	db, err := sql.Open("postgres", app.cfg.dbDsn)
	if err != nil {
		app.errorLog.Println(err)
		return nil, err
	}
	err = db.Ping()

	if err != nil {
		app.errorLog.Println(err)
		return nil, err
	}

	fmt.Println("Database successfully connected")

	return db, nil
}

// Getting variable from .env file
func getVarFromDotenv(key string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return os.Getenv(key)
}

