package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/models"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/middleware"
)

func Signup(user *models.User) (*models.User, error) {
	fmt.Println("Creating a new user in the User Service Layer")
	fmt.Println("--------------------------------------------- \n")

	// Encrypt the password before saving the user
	hashedPassword, err := middleware.HashPassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %v", err)
	}
	user.Password = hashedPassword

	if err := models.Validate(user); err != nil {
		return nil, fmt.Errorf("user validation error: %v", err)
	}

	if err := models.SaveUserToDB(user); err != nil {
		return nil, fmt.Errorf("error saving user to database: %v", err)
	}

	return user, nil
}

func Login(w http.ResponseWriter, user *models.User) (*models.User, error) {
	fmt.Println("Logging in user in the User Service Layer")
	fmt.Println("--------------------------------------------- \n")
	email := *user.Email
	// Retrieve the user from the database
	storedUser, err := models.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user from database: %v", err)
	}

	// Compare the provided password with the stored hashed password
	if !middleware.CheckPasswordHash(user.Password, storedUser.Password) {
		return nil, fmt.Errorf("invalid password")
	}

	// Create a token for the user
	tokenString, err := middleware.CreateToken(*user.Email)
	if err != nil {
		return nil, fmt.Errorf("error creating token: %v", err)
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
		return nil, fmt.Errorf("user ID is required")
	}

	if err := models.SavePortfolioToDB(portfolio); err != nil {
		return nil, fmt.Errorf("error saving portfolio to database: %v", err)
	}

	return portfolio, nil
}

func DeletePortfolio(portfolioID string) error {
	if err := models.DeletePortfolio(portfolioID); err != nil{
		return fmt.Errorf("Error Deleting Portfolio : %v", err)
	}

	return nil
}

func GetPortfolioForUser(userID uuid.UUID) (*models.User, error) {
	user, err := models.GetPortfolioForUser(userID)
	if err != nil {
		return nil, fmt.Errorf("Error Fetching 's' Portfolio")
	}

	return user, nil
}

func VerifyUserAuth(cookie string) (*models.User, error) {
	email, err := middleware.ValidateToken(cookie)
	if err != nil {
		return nil, err
	}

	// also need to check if that particular exists in the DB
	user, err := models.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}