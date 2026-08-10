package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken returns a token with user credentials
func CreateToken(userID uint32) (string, error) {
  permissions := jwt.MapClaims{}
  permissions["authorized"] = true
  permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
  permissions["userID"] = userID

  // secret key
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
  // assign token
  return token.SignedString([]byte(config.SecretKey))
}

// ValidateToken checks the jwt token from header
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, verificationKey)

	if err != nil {
		return err
	}
	
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func verificationKey(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("assign method unexpected! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) != 2 {
		return ""
	}

	return strings.Split(token, " ")[1]
}
