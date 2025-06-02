package main

import (
	"log"
	"net/http"

	"go/hw/4-order-api/configs"
	"go/hw/4-order-api/internal/product"
	"go/hw/4-order-api/middleware"
	"go/hw/4-order-api/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	log.Println(conf)
	db := db.NewDB(conf)
	router := http.NewServeMux()
	productRepository := product.NewProductRepository(db)
	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepository,
	})

	server := http.Server{
		Addr:    ":8082",
		Handler: middleware.JsonLogs(router),
	}

	log.Println("Server is running on port 8082")
	server.ListenAndServe()
}
