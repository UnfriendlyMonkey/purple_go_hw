package verify

import (
	"encoding/json"
	"fmt"
	"go/hw/3-validation-api/configs"
	"go/hw/3-validation-api/pkg/file"
	"go/hw/3-validation-api/pkg/hash"
	"go/hw/3-validation-api/pkg/resp"
	"go/hw/3-validation-api/pkg/send"
	"log"
	"net/http"
)

type VerifyHandlerDeps struct {
	*configs.Config
}

type VerifyHandler struct {
	*configs.Config
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("/verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var addr SendRequest
		err := json.NewDecoder(r.Body).Decode(&addr)
		if err != nil {
			log.Println("no address found in body")
			return
		}
		hashStr, err := hash.Hash(addr.Email)
		if err != nil {
			log.Println("Error getting hash")
			resp.Json(w, "Something's wrong", http.StatusInternalServerError)
		}
		hashedData, err := file.ReadFromFile()
		if err != nil {
			log.Println(err)
			hashedData = make(map[string]string)
		}
		hashedData[hashStr] = addr.Email

		err = file.SaveToFile(hashedData)
		if err != nil {
			fmt.Println("File with emails not found")
			resp.Json(w, "Something's wrong", http.StatusInternalServerError)
		}

		link := fmt.Sprintf("http://localhost:8083/verify/%s", hashStr)
		fmt.Println(link)
		fmt.Println(addr.Email)
		ok, err := send.SendEmail(handler.Config, link, addr.Email)
		var message string
		var status int
		if !ok {
			fmt.Println(err)
			message = fmt.Sprintf("Sending verification email to %s failed", addr.Email)
			status = http.StatusBadRequest
		} else {
			message = fmt.Sprintf("Verification email has been sent to: %s", addr.Email)
			status = http.StatusOK
		}
		data := SendResponse{
			Details: message,
		}
		resp.Json(w, data, status)
		// w.Header().Set("Content-type", "application/json")
		// w.WriteHeader(status)
		// json.NewEncoder(w).Encode(data)
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		hashedData, err := file.ReadFromFile()
		if err != nil {
			fmt.Println("File with emails not found")
			resp.Json(w, "Email not found", http.StatusNotFound)
		}
		addr, exists := hashedData[hash]
		if !exists {
			fmt.Printf("No such hash stored: %s", hash)
			resp.Json(w, "Email not found", http.StatusNotFound)
		}
		delete(hashedData, hash)
		err = file.SaveToFile(hashedData)
		if err != nil {
			log.Println(err)
		}

		data := SendResponse{
			Details: fmt.Sprintf("your address %s is verified", addr),
		}
		resp.Json(w, data, http.StatusOK)
		// w.Header().Set("Content-type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// json.NewEncoder(w).Encode(data)
	}
}
