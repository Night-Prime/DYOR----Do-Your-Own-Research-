package middleware

import (
	"fmt"
	"time"
	"net/http"
	"golang.org/x/crypto/bcrypt"

	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/config"
	"github.com/golang-jwt/jwt/v5"
)

// This handles the authentication and authorization of the user
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateToken(email string) (string, error) {
	cfg := config.Get()
	var secretKey = []byte(cfg.SecretKey)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": 		email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iss":	  	getRole(email),
		"iat":      time.Now().Unix(),
		"aud":      "DYOR",
	})

	token, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	fmt.Println("Token Created: ", token)
	fmt.Println("--------------------------------------------- \n")
	return token, nil
}

func getRole(email string) string {
	if email == "admin"{
		return "admin"
	} else {
		return "user"
	}
}

func ValidateToken(tokenString string) (string, error) {
	cfg := config.Get()
	var secretKey = []byte(cfg.SecretKey)
	
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return claims.Subject, nil
}

func AdminAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("DYOR_token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenString := cookie.Value

		email, err := ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized Admin", http.StatusUnauthorized)
			return
		}
		
		r.Header.Set("email", email)
		next.ServeHTTP(w, r)
	})
}

func UserAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenString := cookie.Value

		email, err := ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized User", http.StatusUnauthorized)
			return
		}
		
		r.Header.Set("email", email)
		next.ServeHTTP(w, r)
	})
}