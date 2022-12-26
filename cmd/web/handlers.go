package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/go-chi/chi"
)

var conn *sql.DB

func (app *application) getAllUsers(w http.ResponseWriter, r *http.Request) {
	var err error
	// u := []user{}
	conn, err = app.createConnection() 
	if err != nil {
		app.errorLog.Println(err)
	}
	err = json.NewEncoder(w).Encode(conn.QueryRow("SELECT * FROM users"))
	if err != nil {
		app.errorLog.Println(err)
		return
	}	
}

func (app *application) registerHand(w http.ResponseWriter, r *http.Request) {
	var err error
	
	
	conn, err = app.createConnection()
	if err != nil {
		app.errorLog.Println(err)
		return
	}
	
	// JSON.Unmarshal(value, pointer)
	
	r.ParseForm()
	
	// Parsing data from forms
	name := r.FormValue("username")
	password := r.FormValue("password")
	date_of_birth := r.FormValue("date_of_birth")
	description := r.FormValue("description")
	address := r.FormValue("address")

	
	ok, _ := checkFormValue(w, r, name, address, date_of_birth, description, password)

	if !ok {
		app.errorLog.Println("Not valid forms value")
		http.Redirect(w, r, "/error", http.StatusFailedDependency)
	}

	qu, err := conn.Query(`
		INSERT 
			INTO users (username, password, address, description, date_ob_birth) 
		VALUES 
			(?, ?, ?, ?, ?)
		`, name, password, address, description, date_of_birth)
	
		
	if err != nil {
		app.errorLog.Println(err)
		return
	}
	defer qu.Close()
	http.Redirect(w, r, "/login", http.StatusOK)
}

func checkFormValue(w http.ResponseWriter, r *http.Request, forms ...string) (bool, string) {
	for _, form := range forms {
		a, _ := regexp.MatchString("^[a-z][A-Z]", r.FormValue(form))
		if r.FormValue(form) == "" {
			return false, "All forms must be fulfilled"
		}
		if a == false {
			return false, "Form must contain only english letters"
		}
	}
	return true, ""
}

func (app *application) getOneUser(w http.ResponseWriter, r *http.Request) {
	var strID string = chi.URLParam(r, "id")

	id, _ := strconv.Atoi(strID)
	
	conn, err := app.createConnection()
	
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	qu, err := conn.Query(
		`SELECT * 
			FROM users WHERE id = ?`, id)
	qu.Close()
	// conn.Close()
}


func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := app.createConnection()
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	r.ParseForm()
	username := r.FormValue("username") 
	password := r.FormValue("password")

	ok, _ := checkFormValue(w, r, username, password)

	if !ok {
		app.errorLog.Println("Not valid forms value")
		http.Redirect(w, r, "/error", http.StatusFailedDependency)
	}

	que, err := conn.Query(`
		INSERT * FROM users WHERE username = ? AND password = ?
	`, username, password)

	defer que.Close()
}

func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {
	conn, err := app.createConnection()
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	r.ParseForm()
	username := r.FormValue("username") 
	password := r.FormValue("password")

	ok, _ := checkFormValue(w, r, username, password)

	if !ok {
		app.errorLog.Println("Not valid forms value")
		http.Redirect(w, r, "/error", http.StatusFailedDependency)
	}

	que, err := conn.Query(`
		DELETE * FROM users WHERE username = ? AND password = ?
	`, username, password)

	defer que.Close()
}

func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {
	conn, err := app.createConnection()
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	r.ParseForm()
	username := r.FormValue("username") 
	password := r.FormValue("password")

	newUsername := r.FormValue("newUsername")
	newAddress := r.FormValue("newAdderss")
	newPassword := r.FormValue("newPassword")
	newDescription := r.FormValue("newDescription")
	newDoB := r.FormValue("newDoB")

	ok, _ := checkFormValue(w, r, username, password)

	if !ok {
		app.errorLog.Println("Not valid forms value")
		http.Redirect(w, r, "/error", http.StatusFailedDependency)
	}

	que, err := conn.Query(`
		UPDATE 
			users (username, password, address, description, date_of_birth)
		SET 
			username = ?
			password = ?
			address = ? 
			description = ?
			date_of_birth = ?
		WHERE 
			username = ? AND password = ?
	`, newUsername, newPassword, newAddress, newDescription, newDoB, username, password)
	if err != nil {
		app.errorLog.Println(err)
		return
	}
	
	defer que.Close()
}
