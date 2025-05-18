package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func randomHandler(w http.ResponseWriter, r *http.Request) {
	num := rand.Intn(6) + 1
	w.Write(fmt.Appendf(nil, "%d", num))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/random", randomHandler)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
