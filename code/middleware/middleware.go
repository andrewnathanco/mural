package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/labstack/echo/v4"
)

var (
	ErrCouldNotGetGameKey = fmt.Errorf("could not get game key")
)

func GetGameKeyFromContext(c echo.Context) (string,error) {
	game_key, ok := c.Get("game-key").(string)
	if !ok {
		return "", ErrCouldNotGetGameKey
	}

	if game_key == "" {
		return "", ErrCouldNotGetGameKey
	}

	return game_key, nil
} 

// Process is the middleware function.
func HashRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user_agent := c.Request().UserAgent()
		game_key := HashString(user_agent)

		c.Set("game-key", game_key)
		return next(c)
	}
}

func HashString(data string) string {
	// Calculate the SHA-256 hash
	hasher := sha256.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

