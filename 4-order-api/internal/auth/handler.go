package auth

import (
	"go/hw/3-validation-api/pkg/resp"
	"go/hw/4-order-api/configs"
	"go/hw/4-order-api/pkg/req"
	"net/http"
)

type AuthHandlerDeps struct {
	AuthService *AuthService
	Config      *configs.Config
}

type AuthHandler struct {
	*AuthService
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) *AuthHandler {
	handler := &AuthHandler{AuthService: deps.AuthService, Config: deps.Config}
	router.HandleFunc("POST /auth/send-code", handler.SendCode())
	router.HandleFunc("POST /auth/verify-code", handler.VerifyCode())
	return handler
}

func (h *AuthHandler) SendCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[SendCodeRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data, err := h.AuthService.SendCode(body.Phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp.Json(w, data, http.StatusOK)
	}
}

func (h *AuthHandler) VerifyCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[VerifyCodeRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		token, err := h.AuthService.VerifyCode(body.SessionID, body.Code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		data := VerifyCodeResponse{
			Token: token,
		}
		resp.Json(w, data, http.StatusOK)
	}
}
