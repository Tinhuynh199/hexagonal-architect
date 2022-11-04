package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	. "hexrestapi1/internal/infrastructure/domain/user"
	. "hexrestapi1/internal/infrastructure/service/user"
)

const (
	InternalServerError = "Internal Server Error"
)

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetAllUsers implements port.UserTransport
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, users)
}

// GetUser implements port.UserTransport
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, user)
}

// CreateUser implements port.UserTransport
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	er1 := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if er1 != nil {
		
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}
	
	res, er2 := h.service.CreateUser(r.Context(), &user)
	
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusCreated, res)
}

// UpdateUser implements port.UserTransport
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
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

	res, er2 := h.service.UpdateUser(r.Context(), &user)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, res)
}

// DeleteUser implements port.UserTransport
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	res, err := h.service.DeleteUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, res)
}

func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}
