package handlers

import(
	"time"
	"net/http"
	"encoding/json"
    "fmt"
	
	"github.com/google/uuid"
    "github.com/lib/pq"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/models"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/service"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/errors"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	
	loggedInUser, err := service.Login(w, user)
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
		switch err.(type) {
        case *errors.ValidationError:
            http.Error(w, err.Error(), http.StatusBadRequest) 
        case *errors.DatabaseError:
            http.Error(w, err.Error(), http.StatusInternalServerError) 
		case *errors.CustomError:
            http.Error(w, err.Error(), http.StatusInternalServerError)
        default:
            http.Error(w, err.Error(), http.StatusInternalServerError) 
        }
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
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	parsedID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	user, err := service.GetPortfolioForUser(parsedID)
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
	json.NewEncoder(w).Encode(user)
}

// this code is weird.
func UpdatePortfolioHandler(w http.ResponseWriter, r *http.Request) {
    // TODO : Convert to reusable Model
    var request struct {
        Portfolio struct {
            ID              uuid.UUID  `json:"id"`
            Name            string     `json:"name"`
			UserID          uuid.UUID  `json:"user_id"`
            TotalValue      float64    `json:"total_value,omitempty"`
            InvestmentGoals []string   `json:"investment_goals,omitempty"`
            AssetPreference []string   `json:"asset_preference,omitempty"`
        } `json:"portfolio"`
        Assets struct {
            Add    []struct {
                Symbol       string    `json:"symbol"`
                Name         string    `json:"name"`
                Type         string    `json:"type"`
                Quantity     float64   `json:"quantity,omitempty"`
                CurrentPrice float64   `json:"current_price,omitempty"`
                Volume       float64   `json:"volume,omitempty"`
            } `json:"add"`
            Update []struct {
                ID           uuid.UUID `json:"id"`
                Symbol       string    `json:"symbol"`
                Name         string    `json:"name"`
                Quantity     float64   `json:"quantity,omitempty"`
                CurrentPrice float64   `json:"current_price,omitempty"`
                Volume       float64   `json:"volume,omitempty"`
            } `json:"update"`
            Delete []uuid.UUID `json:"delete"`
        } `json:"assets"`
    }

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        fmt.Printf("JSON decode error: %v", err) 
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Prepare portfolio update
    portfolio := &models.Portfolio{
        ID:              request.Portfolio.ID,
        UserID:          request.Portfolio.UserID,
        Name:            request.Portfolio.Name,
        TotalValue:      request.Portfolio.TotalValue,
        InvestmentGoals: pq.StringArray(request.Portfolio.InvestmentGoals),
        AssetPreference: pq.StringArray(request.Portfolio.AssetPreference),
    }

    var assetsToAdd []*models.Asset
    for _, asset := range request.Assets.Add {
        assetType := models.AssetType(asset.Type)
        if assetType != models.AssetTypeStock && assetType != models.AssetTypeCrypto {
            http.Error(w, "Invalid asset type", http.StatusBadRequest)
            return
        }

        assetsToAdd = append(assetsToAdd, &models.Asset{
            AssetBase: models.AssetBase{
                Type:         assetType,
                Symbol:       asset.Symbol,
                Name:         asset.Name,
                Quantity:     asset.Quantity,
                CurrentPrice: asset.CurrentPrice,
                Volume:       asset.Volume,
            },
        })
    }

    var assetsToUpdate []*models.Asset
    for _, asset := range request.Assets.Update {
        assetsToUpdate = append(assetsToUpdate, &models.Asset{
            AssetBase: models.AssetBase{
                ID:           asset.ID,
                Symbol:       asset.Symbol,
                Name:         asset.Name,
                Quantity:     asset.Quantity,
                CurrentPrice: asset.CurrentPrice,
                Volume:       asset.Volume,
            },
        })
    }

    updateReq := &models.PortfolioUpdateRequest{
        Portfolio:      portfolio,
        AssetsToAdd:    assetsToAdd,
        AssetsToUpdate: assetsToUpdate,
        AssetsToDelete: request.Assets.Delete,
    }

    // Update portfolio
    updatedPortfolio, err := service.UpdateUserPortfolio(updateReq)
    if err != nil {
        switch err.(type) {
        case *errors.ValidationError:
            http.Error(w, err.Error(), http.StatusBadRequest)
        case *errors.DatabaseError:
            http.Error(w, err.Error(), http.StatusInternalServerError)
        default:
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(updatedPortfolio)
}


func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
        Name:     "token",
        Value:    "",
        Path:     "/",
        Expires:  time.Unix(0, 0),
        HttpOnly: true,
        Secure:   true,
        SameSite: http.SameSiteStrictMode,
    })
    
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Successfully logged out"})
}