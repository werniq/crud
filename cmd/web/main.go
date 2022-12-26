package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

type config struct {
	port 	   int
	dbDsn 	   string
	enviroment string
}

type user struct{
	id 			int 		`json:"id"`
	name 	    string 		`json:"name"`
	address 	string 		`json:"address"`
	dateOfBirth time.Time 	`json:"date_of_birth"` 
	createdAt 	time.Time 	`json:"-"`
}


type application struct {
	cfg 		config
	infoLog 	*log.Logger
	errorLog 	*log.Logger
	db 			Database
}

func (app *application) serve() error {
	server := &http.Server{
		Addr: fmt.Sprintf(":%d",app.cfg.port),
		IdleTimeout: 30 * time.Second,
		ReadTimeout: 10*time.Second,
	}

	return server.ListenAndServe()
}

func main() {	
	var cfg config
	var app application

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	conn, err := app.createConnection()
	db := Database{DB: conn}

	if err != nil {
		errorLog.Println(err)
	}
	err = conn.Ping()

	if err != nil{
		errorLog.Println(err)
	}

	app = application{
		cfg: cfg,
		infoLog: infoLog,
		errorLog: errorLog,
		db: db,
	}

	err = app.serve()
	if err != nil {
		errorLog.Println(err)
	}
}

