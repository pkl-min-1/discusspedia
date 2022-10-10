package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthErrorResponse struct {
	Error string `json:"error"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, AuthErrorResponse{Error: "No token provided"})
			return
		}

		tokenString = tokenString[(len("Bearer ")):]

		token, err := ValidateToken(tokenString)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, AuthErrorResponse{Error: "Invalid token"})
				return
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, AuthErrorResponse{Error: "Bad Request"})
			return
		}

		if token.Valid {
			// claims := token.Claims.(*Claims)
			// fmt.Println(claims.Email)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, AuthErrorResponse{Error: "Invalid token"})
		}

	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		fmt.Println(c.Request.Method)

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

