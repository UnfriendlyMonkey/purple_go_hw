package main

import (
	"fmt"
	"go/hw/3-validation-api/configs"
	"go/hw/3-validation-api/verify"
	"log"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	fmt.Println("Here I am")
	fmt.Println(conf)
	router := http.NewServeMux()
	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{
		Config: conf,
	})
	log.Println(router)
	server := http.Server{
		Addr:    ":8083",
		Handler: router,
	}
	log.Println("Server is listening on port 8083")
	server.ListenAndServe()
}
