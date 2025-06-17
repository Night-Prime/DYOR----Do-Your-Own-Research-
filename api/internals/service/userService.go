package service

import (
	"fmt"
	"net/http"
	"time"
	"os"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/models"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/middleware"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/errors"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/database"
)

// TODO Social Login & Password Recovery feature
func Signup(user *models.User) (*models.User, error) {
	fmt.Println("Creating a new user in the User Service Layer")
	fmt.Println("--------------------------------------------- \n")

	hashedPassword, err := middleware.HashPassword(user.Password)
	if err != nil {
        return nil, &errors.DatabaseError{
            Message: "Error hashing password",
            Err: err,
        }
    }
	user.Password = hashedPassword

	if err := models.Validate(user); err != nil {
		return nil, err
	}

	if err := models.SaveUserToDB(user); err != nil {
		return nil, err
	}

	return user, nil
}

func Login(w http.ResponseWriter, user *models.User) (*models.User, error) {
	fmt.Println("Logging in user in the User Service Layer")
	fmt.Println("--------------------------------------------- \n")
	email := *user.Email

	storedUser, err := models.GetUserByEmail(email)
	if err != nil {
		fmt.Printf("Error while saving User: %v", err)
		return nil, err
	}

	if !middleware.CheckPasswordHash(user.Password, storedUser.Password) {
		return nil, &errors.ValidationError{Message:"Invalid password"}
	}

	tokenString, err := middleware.CreateToken(*user.Email)
	if err != nil {
		return nil,  &errors.DatabaseError{
            Message:"Error while creating token",
			Err: err,
		}
	}

	// Set the token as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	return storedUser, nil
}

func CreatePortfolio(portfolio *models.Portfolio) (*models.Portfolio, error) {
	fmt.Println("Creating a new portfolio for the User in the User Service Layer")
	fmt.Println("--------------------------------------------- \n")

	if portfolio.UserID.String() == "" {
		return nil, &errors.ValidationError{Message:"user ID is required"}
	}

	if err := models.SavePortfolioToDB(portfolio); err != nil {
		return nil, err
	}

	return portfolio, nil
}

func DeletePortfolio(portfolioID string) error {
	if err := models.DeletePortfolio(portfolioID); err != nil{
		return err
	}

	return nil
}

func GetPortfolioForUser(userID uuid.UUID) (*models.User, error) {
	user, err := models.GetPortfolioForUser(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func VerifyUserAuth(cookie string) (*models.User, error) {
	email, err := middleware.ValidateToken(cookie)
	if err != nil {
		return nil, err
	}

	user, err := models.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}


func UpdateUserPortfolio(req *models.PortfolioUpdateRequest) (*models.Portfolio, error) {
    fmt.Println("Updating portfolio and assets in the Portfolio Service Layer")
    fmt.Println("----------------------------------------------------------")

    if req.Portfolio.ID == uuid.Nil {
        return nil, &errors.ValidationError{Message: "Portfolio ID is required"}
    }
	// TODO : Optimize Query & refactor

    if err := models.UpdatePortfolio(
        req.Portfolio,
        req.AssetsToAdd,
        req.AssetsToUpdate,
        req.AssetsToDelete,
    ); err != nil {
        return nil, err
    }

    // Refetching the updated portfolio (logic is debatable)
    var updatedPortfolio models.Portfolio
    if err := database.GetDB().Preload("Assets").
        First(&updatedPortfolio, "id = ?", req.Portfolio.ID).Error; err != nil {
        return nil, &errors.DatabaseError{
            Message: "Error fetching updated portfolio", 
            Err: err,
        }
    }

    return &updatedPortfolio, nil
}

func GetJSONFile(filePath string) (map[string]interface{}, error) {
	fmt.Println("Grabbing the JSON Data")
    fmt.Println("----------------------------------------------------------")

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, &errors.DatabaseError{
			Message: "Unable to load data",
			Err: err,
		}
	}

	// parsing occurs:
	var data map[string]interface{}
	if err := json.Unmarshal(fileContent, &data); err != nil {
		return nil, &errors.DatabaseError{
			Message: "Unable to parse data",
			Err: err,
		}
	}

	return data, nil
}