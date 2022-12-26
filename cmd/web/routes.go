package main

import (
	"github.com/go-chi/chi"
)

func (app *application) handlers() {
	s := chi.NewRouter()
	
	s.Get("/users", app.getAllUsers)
	s.Get("/users/{id}", app.getOneUser)
	s.Post("/register", app.registerHand)
	s.Post("/login", app.loginHandler)
	s.Delete("/users/{id}", app.deleteUser)
	s.Put("/users/{id}", app.updateUser)


}