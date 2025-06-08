package main

import (
	"log"
	"net/http"

	"go/hw/4-order-api/configs"
	"go/hw/4-order-api/internal/auth"
	"go/hw/4-order-api/internal/product"
	"go/hw/4-order-api/internal/user"
	"go/hw/4-order-api/middleware"
	"go/hw/4-order-api/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	log.Println(conf)
	db := db.NewDB(conf)
	router := http.NewServeMux()

	// repositories
	productRepository := product.NewProductRepository(db)
	userRepository := user.NewUserRepository(db)

	// services
	authService := auth.NewAuthService(userRepository)

	// handlers
	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepository,
		Config:            conf,
	})
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		AuthService: authService,
		Config:      conf,
	})

	server := http.Server{
		Addr:    ":8082",
		Handler: middleware.JsonLogs(router),
	}

	log.Println("Server is running on port 8082")
	server.ListenAndServe()
}
