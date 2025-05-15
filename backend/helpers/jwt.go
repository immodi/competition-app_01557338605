package helpers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("25dfb0830d20d65fe62c34b4755f4693d1969add26ba74a2b4b9e3f64e87c9959f72e50657e63a17c5c68dd6ce115a9d1257266be5f7d3cbd027a6286b9dea2dc01f6e5cfb1f50ab703f9a5ed36bc1d8c92e6178a86900316b22a3f328d21b892116d20d937f1a1abcce335a2637f28cde6d2ed7e392607f79c77d3981bcccb9644ff7fb8305a65d67a429ef17c88696a68962c0d94cb388cd4203676440de4fb49dcf37febb0010f17e669cece5d7999489c3279cfef3cd373a8f544ab0e5df068904f9db87d7d060e8e2208fe760be8fcddcd2e8d1da764a7ceea7fc8a7087aabb5d0e97e3c878773946bf156f2ce1f4e567dee0a3109d96d17398c5482039")

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
		HttpError(w, http.StatusUnauthorized, "no token?")
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
