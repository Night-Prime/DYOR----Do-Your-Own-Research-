package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/models"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/service"
	"github.com/google/uuid"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	loggedInUser, err := service.Login(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(loggedInUser)
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdUser, err := service.Signup(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func CreatePortfolioHandler(w http.ResponseWriter, r *http.Request) {
	portfolio := &models.Portfolio{}

	if err := json.NewDecoder(r.Body).Decode(portfolio); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userPortfolio, err := service.CreatePortfolio(portfolio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userPortfolio)
}

func DeletePortfolioHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	err := service.DeletePortfolio(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetPortfolioForUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	parsedID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid ID format, must be a valid UUID", http.StatusBadRequest)
		return
	}

	userPortfolio, err := service.GetPortfolioForUser(parsedID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userPortfolio)
}