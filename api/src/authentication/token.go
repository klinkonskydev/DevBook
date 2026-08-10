package authentication

import (
	"api/src/config"
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
