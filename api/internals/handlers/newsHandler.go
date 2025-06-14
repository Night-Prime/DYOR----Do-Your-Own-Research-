package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/errors"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/service"
)

type NewsHandler struct {
	newsService *service.NewsService
}

func NewNewsHandler(newsService *service.NewsService) *NewsHandler {
	return &NewsHandler{
		newsService: newsService,
	}
}

func (h *NewsHandler) GetNewsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	news, err := h.newsService.GetNews()
	if err != nil {
		switch err.(type) {
		case *errors.DatabaseError:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	response, err := json.Marshal(news)
	if err != nil {
		http.Error(w, "Error marshalling response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *NewsHandler) CalculateNewsSentimentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sentimentScore, err := h.newsService.CalculateNewsSentiment()
	if err != nil {
		switch err.(type) {
		case *errors.DatabaseError:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		case *errors.ValidationError:
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	response := map[string]float64{"sentiment_score": sentimentScore}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error marshalling response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (h *NewsHandler) GetTopGainersLosersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	topGainersLosers, err := h.newsService.GetTopGainersLosers()
	if err != nil {
		switch err.(type) {
		case *errors.DatabaseError:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	response, err := json.Marshal(topGainersLosers)
	if err != nil {
		http.Error(w, "Error marshalling response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}