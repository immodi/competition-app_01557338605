package helpers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("failed to extract claims")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("username claim not found or not a string")
	}

	return username, nil
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request, isQualifiedCallback func(username string) bool, handler func(w http.ResponseWriter, r *http.Request)) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		HttpError(w, http.StatusUnauthorized, "request does not contain an access token")
		return
	}
	tokenString = tokenString[len("Bearer "):]

	username, err := verifyToken(tokenString)
	if err != nil {
		HttpError(w, http.StatusUnauthorized, "invalid token, you dont have permission for this route")
		return
	}

	if isQualifiedCallback != nil && !isQualifiedCallback(username) {
		HttpError(w, http.StatusUnauthorized, "you dont have permission for this route")
		return
	}

	handler(w, r)
}
