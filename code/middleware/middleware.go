package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

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
func GenerateUserKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user_agent := c.Request().UserAgent()
		user_key := HashString(user_agent)

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

