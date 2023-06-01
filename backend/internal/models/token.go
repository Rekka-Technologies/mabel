package models

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GenerateToken generates a jwt token for a given user, this utilises multiple
// server env variables to determine the token lifespan and the secret key.
// The token is signed using HMAC and SHA256.
func GenerateToken(user User) (string, error) {

	// Fetch the Lifespan of the token from the env file, this will be used
	// to determine when the token will expire.
	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		return "", err
	}

	// Build the token
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user.ID
	claims["user_name"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	return token.SignedString([]byte(os.Getenv("API_SECRET")))

}

// TokenValid checks if the token is valid or not, if it is valid then it
// returns the token object otherwise it returns nil and an error.
func TokenValid(c *gin.Context) (*jwt.Token, error) {
	// Get the token from the Authorization header
	rawbearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(rawbearerToken, " ")) != 2 {
		return nil, fmt.Errorf("invalid token")
	}

	// Parse the token and validate it
	bearerToken := strings.Split(rawbearerToken, " ")[1]
	token, err := jwt.Parse(bearerToken,
		func(token *jwt.Token) (interface{}, error) {
			// We only allow HMAC signing
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("API_SECRET")), nil
		},
	)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// ExtractTokenUID extracts the user_id from the token after it has the token
// has been validated so we can use it in determining what data to return.
func ExtractTokenUID(c *gin.Context) (uint, error) {
	token, err := TokenValid(c)
	if err != nil {
		return 0, err
	}

	// If the token is valid, extract the user_id from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}

	return 0, nil
}
