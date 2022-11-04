package app

import (
	"log"
	"net/http"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

// // Set all required routers
func (a *App) setRouters() {

	// Routing for handling the project of manipulate User Info
	userPath := "/users"
	a.Get(userPath, a.User.GetAllUsers)
	a.Get(userPath+"/{id}", a.User.GetUser)
	a.Post(userPath, a.User.CreateUser)
	a.Put(userPath+"/{id}", a.User.UpdateUser)
	a.Delete(userPath+"/{id}", a.User.DeleteUser)

	// Routing for handling the project of manipulate Student Info
	studentPath := "/students"
	a.Get(studentPath, a.Student.GetAllStudents)
	a.Get(studentPath+"/{id}", a.Student.GetStudent)
	a.Post(studentPath, a.Student.CreateStudent)
	a.Put(studentPath+"/{id}", a.Student.UpdateStudent)
	a.Delete(studentPath+"/{id}", a.Student.DeleteStudent)

}

// // Wrap  the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.StrictSlash(true).HandleFunc(path, f).Methods(GET)
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods(POST)
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods(PUT)
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods(DELETE)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
