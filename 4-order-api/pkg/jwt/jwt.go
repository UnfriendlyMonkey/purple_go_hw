package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
	Phone     string `json:"phone"`
}

type JWT struct {
	SecretKey string
}

func NewJWT(secretKey string) *JWT {
	return &JWT{SecretKey: secretKey}
}

func (j *JWT) GenerateToken(data *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    data.UserID,
		"session_id": data.SessionID,
		"phone":      data.Phone,
	})
	tokenString, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWT) VerifyToken(tokenString string) (bool, *Claims) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(j.SecretKey), nil
	})
	if err != nil {
		return false, nil
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, nil
	}
	return token.Valid, &Claims{
		UserID:    claims["user_id"].(string),
		SessionID: claims["session_id"].(string),
		Phone:     claims["phone"].(string),
	}
}
