package auth

import (
	"strings"
	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"status": 401, "message": message})
}

func AuthMiddleware() gin.HandlerFunc {
	return func (c *gin.Context) {
		authHeader := c.Request.Header["Authorization"]
		authorized := false
	
		if len(authHeader) > 0 {
			tokenSlice := authHeader[0]
	
			if tokenSlice != "" {
				apiAccessString := strings.Split(tokenSlice, " ")[1]
	
				if strings.Contains(apiAccessString, ":") {
					apiAccess := strings.Split(apiAccessString, ":")
	
					if len(apiAccess) == 2 {
						APIKey := apiAccess[0]
						APISecret := apiAccess[1]
			
						for _, user := range users {
							if user.APIKey == APIKey && user.APISecret == APISecret {
								authorized = true
								break
							}
						}
					}
				}
			}
		}
	
		if authorized == false {
			respondWithError(c, 401, "Unauthorized")
			return
		}

		c.Next()
	}
}