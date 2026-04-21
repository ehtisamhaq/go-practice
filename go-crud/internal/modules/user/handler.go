package user

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *service
}

func NewHandler(service *service) *Handler {
	return &Handler{service: service}
}

// Handler methods for Create, GetAll, GetByID, Update, Delete would go here
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.service.Create(u)
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetByID(id int) (User, error) {
	return h.service.GetByID(id)
}

func (h *Handler) Update(user User) error {
	return h.service.Update(user)
}

func (h *Handler) Delete(id int) error {
	return h.service.Delete(id)
}
