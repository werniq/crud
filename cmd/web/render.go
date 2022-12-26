package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, page string) error {
	var t *template.Template
	var err error 
	ttr := fmt.Sprintf("templates/%s.gohtml", page)
	t, err = app.parseTempalte(w, r, ttr)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}
	err = t.Execute(w, nil)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}
	return nil
}

func (app *application) parseTempalte(w http.ResponseWriter, r *http.Request, page string) (*template.Template, error) {
	var tmp *template.Template
	var e error 

	// tmp = template.New(page)	
	tmp, e = template.ParseGlob(page)
	if e != nil {
		app.errorLog.Println(e)
		return nil, e
	}

	e = tmp.Execute(w, nil)
	
	if e != nil {
		app.errorLog.Println(e)
		return nil, e
	}
	
	return tmp, nil
}