package middlewares

import (
	"github.com/labstack/echo/v4"
)

// SetAuthContext sets user struct in the echo context as part of the middleware process
func SetAuthContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			// 1. get the string of JWT from header

			// 2. validate the token

			// 3. set in the echo context and return
			return next(c)
		}
	}
}

// (If there is a login function in the usecase layer, we can bring this function to the inner side)
// validateToken validates the token string, and returns the ID.
func validateToken(tokenString string) (id string, err error) {

	// 1. validate the token

	// 2. retrieve id from the token

	return id, nil
}
