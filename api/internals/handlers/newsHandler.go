package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/errors"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/service"
)

func GetNewsHandler(w http.ResponseWriter, r *http.Request) {
	news, err := service.GetNews()
	if err != nil {
		switch err.(type) {
		case *errors.ValidationError:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case *errors.DatabaseError:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(news); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}