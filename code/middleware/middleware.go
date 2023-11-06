package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var (
	ErrCouldNotGetUserKey = fmt.Errorf("could not get user key")
)

func GetUserKeyFromContext(c echo.Context) (string,error) {
	user_key, ok := c.Get("user-key").(string)
	if !ok {
		return "", ErrCouldNotGetUserKey
	}

	if user_key == "" {
		return "", ErrCouldNotGetUserKey
	}

	return user_key, nil
} 

// Process is the middleware function.
func GetUserKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user_key_cook, err := c.Cookie("user-key")
		var user_key string
		if err != nil {
			user_key = fmt.Sprintf("%v", uuid.New())  
		} else {
			user_key = user_key_cook.Value
		}

		cookie := new(http.Cookie)
		cookie.Name = "user-key"
		cookie.Value = user_key
		c.SetCookie(cookie)
		c.Set("user-key", user_key)
		return next(c)
	}
}

func HashString(data string) string {
	// Calculate the SHA-256 hash
	hasher := sha256.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

