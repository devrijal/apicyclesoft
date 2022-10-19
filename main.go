package main

import (
	"api-cyclesoft/customer"
	"api-cyclesoft/auth"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type routes struct {
	router *gin.Engine
}

func NewRoutes() routes {
	r := routes{
		router: gin.Default(),
	}

	r.router.SetTrustedProxies([]string{"127.0.0.1"})
	v1 := r.router.Group("/v1")
	v1.Use(auth.AuthMiddleware())

	r.CustomerRoutes(v1)
	return r
}

func (r routes) CustomerRoutes(rg *gin.RouterGroup) {
	customers := rg.Group("/customers")

	customers.GET("/", customer.GetCustomers)
	customers.POST("/", customer.CreateCustomer)

	customers.GET("/:id", customer.GetCustomerByID)
	customers.PUT("/:id", customer.UpdateCustomer)
}

func main() {

	envErr := godotenv.Load()
	if envErr != nil {
		fmt.Printf("Error loading credentials: %v", envErr)
	}

	NewRoutes().router.Run(fmt.Sprintf(":%s", os.Getenv("GIN_PORT")))
	gin.SetMode(gin.ReleaseMode)
}
