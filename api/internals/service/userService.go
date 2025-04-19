package service

import (
	"fmt"
	"net/http"
	"time"
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

func UserLogin(w http.ResponseWriter, user *models.User) (*models.User, error) {
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
		Name:     "DYOR_token",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	return storedUser, nil
}