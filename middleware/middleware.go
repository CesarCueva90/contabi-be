package middleware

import (
	"contabi-be/usecase"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	UseCase usecase.LoginUseCase
}

func New(useCase usecase.LoginUseCase) *Middleware {
	return &Middleware{UseCase: useCase}
}

// CORS configures Cross-Origin Resource Sharing middleware
func (m *Middleware) CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},  // Allow all origins
		AllowMethods:     []string{"*"},  // Allow all HTTP methods
		AllowHeaders:     []string{"*"},  // Allow all headers
		ExposeHeaders:    []string{"*"},  // Expose all headers
		AllowCredentials: true,           // Allow credentials
		MaxAge:           12 * time.Hour, // Cache preflight requests for 12 hours
	})
}

// AuthMiddleware validates login credentials passed via request headers
func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve credentials from headers
		username := c.GetHeader("X-Username")
		password := c.GetHeader("X-UserPassword")

		// Check if both username and password are provided
		if username == "" || password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required in headers"})
			c.Abort()
			return
		}

		// Call UseCase to validate credentials
		user, err := m.UseCase.Login(username, password)
		if err != nil || user.ID == "" {
			// If credentials are invalid, return error
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			c.Abort()
			return
		}

		// Continue with the next middleware or controller
		c.Next()
	}
}
