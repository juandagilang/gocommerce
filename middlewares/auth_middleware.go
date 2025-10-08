package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("a038fd667db4d281564f9729cf2e86c4972f48f81df0219df124c71143123ddc") // Samakan dengan create_token_handler.go

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		if tokenString == "" {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"message": "Unauthorized", "error": err.Error(), "token": tokenString})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["user_id"].(float64))

		// Print userID dan token untuk debug
		println("DEBUG userID:", userID)
		println("DEBUG token:", tokenString)

		c.Set("user_id", userID)
		c.Next()
	}
}
