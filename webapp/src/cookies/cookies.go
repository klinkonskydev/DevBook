package cookies

import (
	"net/http"
	"time"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Setup uses the environment variables to create the SecureCookie
func Setup() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Save stores the authentication data
func Save(w http.ResponseWriter, ID, token string) error {
	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	encodedData, err := s.Encode("auth", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    encodedData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

// Read returns the values stored in the cookie
func Read(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("auth")
	if err != nil {
		return nil, err
	}

	values := make(map[string]string)
	if err = s.Decode("auth", cookie.Value, &values); err != nil {
		return nil, err
	}

	return values, nil
}

// Delete removes the values stored in the cookie
func Delete(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
