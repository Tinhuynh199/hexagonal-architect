package port

import (
	"net/http"
)

// Driver Actor -- App -> Core
type StudentTransport interface {
	GetAllStudents(w http.ResponseWriter, r *http.Request)
	GetStudent(w http.ResponseWriter, r *http.Request)
	CreateStudent(w http.ResponseWriter, r *http.Request)
	UpdateStudent(w http.ResponseWriter, r *http.Request)
	DeleteStudent(w http.ResponseWriter, r *http.Request)
}