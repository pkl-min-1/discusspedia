package api

import (
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

// func AllowOrigin (c *gin.Context) {
// 	gin.Default().RedirectTrailingSlash = true

// 	c.Header("Access-Control-Allow-Origin", "*")
// 	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE")
// 	c.Header("Access-Control-Allow-Headers", "*")
// 	c.Header("Access-Control-Allow-Credentials", "true")
// 	c.Header("Content-Type", "application/json")
// }
