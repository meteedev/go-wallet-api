package apperrs

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}


func CustomErrorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c) // Call the next handler
		if err != nil {
			code := http.StatusInternalServerError
			message := "Internal Server Error"

			// Check for specific error types and customize error response
			switch e := err.(type) {
			case *echo.HTTPError:
				code = e.Code
				message = e.Message.(string)
			default:
				// Log the error for debugging purposes
				log.Println(err)
			}

			// Return standardized JSON response with error details
			return c.JSON(code, CustomError{Code: code, Message: message})
		}
		return nil
	}
}