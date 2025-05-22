package handlers

import(
	"net/http"
	"encoding/json"
	
	"github.com/google/uuid"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/models"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/service"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	
	loggedInUser, err := service.Login(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
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

	// handle portfolio creation:
    userPortfolio := &models.Portfolio{
        UserID:      createdUser.ID,
        Name:        createdUser.FirstName + "'s Portfolio",
        TotalValue:  0,
    }

    _, err = service.CreatePortfolio(userPortfolio)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdUser)
}

// auth_handlers.go
func VerifyUser(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("token")
    if err != nil {
        http.Error(w, "No session token", http.StatusUnauthorized)
        return
    }

    user, err := service.VerifyUserAuth(cookie.Value)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
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
	json.NewEncoder(w).Encode(userPortfolio)
}

func DeletePortfolioHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	err := service.DeletePortfolio(id) 
	if err != nil{
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

	// Validate if the ID is a valid UUID
	if _, err := uuid.Parse(id); err != nil {
		http.Error(w, "Invalid ID format, must be a valid UUID", http.StatusBadRequest)
		return
	}

	parsedID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid ID format, must be a valid UUID", http.StatusBadRequest)
		return
	}

	user, err := service.GetPortfolioForUser(parsedID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}