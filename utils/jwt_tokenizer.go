package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getJWTSecret() string {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		fmt.Println("Not able to fetch jwt secret")
	}

	return jwtSecret
}

func GenerateJWT(userID string) (string, error) {
	jwtSecret := getJWTSecret()

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecret))
}

func ValidateJWT(tokenString string) (string, error) {
	jwtSecret := getJWTSecret()

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure token is signed with correct algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(jwtSecret), nil
	})

	// Handle parsing error
	if err != nil {
		return "", errors.New("Invalid or expired token")
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("Invalid token claims")
	}

	// Get userId from token
	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("userID missing in token")
	}

	return userID, nil
}
