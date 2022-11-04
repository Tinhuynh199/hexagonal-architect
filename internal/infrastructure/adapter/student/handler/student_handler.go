package handler

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"

	. "hexrestapi1/internal/infrastructure/domain/student"
	. "hexrestapi1/internal/infrastructure/service/student"
)

type StudentHandler struct {
	service StudentService
}

func NewStudentHandler(service StudentService) *StudentHandler {
	return &StudentHandler{service: service}
}

// GetAllStudents implements port.StudentTransport
func (h *StudentHandler) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students, err := h.service.GetAllStudents(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, students)
}

// GetStudent implements port.StudentTransport
func (h *StudentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetStudent(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusOK, user)
}

// CreateStudent implements port.StudentTransport
func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var user Student
	er1 := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}

	res, er2 := h.service.CreateStudent(r.Context(), &user)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusOK, res)
}

// UpdateStudent implements port.StudentTransport
func (h *StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var user Student
	er1 := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}

	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	if len(user.ID) == 0 {
		user.ID = id
	} else if id != user.ID {
		http.Error(w, "Id not match", http.StatusBadRequest)
		return
	}

	res, er2 := h.service.UpdateStudent(r.Context(), &user)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusOK, res)
}

// DeleteStudent implements port.StudentTransport
func (h *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	res, err := h.service.DeleteStudent(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if res == 1 {
		res1 := "The record is successful deleted !!!"
		JSON(w, http.StatusOK, res1)
	}
}

// GetAllStudent implements port.StudentTransport
func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}