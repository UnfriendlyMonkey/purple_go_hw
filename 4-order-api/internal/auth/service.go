package auth

import (
	"errors"
	"go/hw/4-order-api/internal/user"
	"go/hw/4-order-api/pkg/authutils"
	"go/hw/4-order-api/pkg/jwt"
	"os"
	"strconv"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (s *AuthService) SendCode(phone string) (*SendCodeResponse, error) {
	code := authutils.GenerateCode()
	sessionID := authutils.GenerateSessionID()
	_, err := authutils.SendCode(phone, code)
	if err != nil {
		return nil, err
	}
	existingUser, _ := s.UserRepository.GetUserByPhone(phone)
	if existingUser != nil {
		existingUser.SessionID = sessionID
		existingUser.Code = code
		_, err := s.UserRepository.UpdateUser(existingUser)
		if err != nil {
			return nil, err
		}
	} else {
		user := &user.User{
			Phone:     phone,
			Code:      code,
			SessionID: sessionID,
		}
		_, err = s.UserRepository.CreateUser(user)
		if err != nil {
			return nil, err
		}
	}
	return &SendCodeResponse{
		Code:      code,
		SessionID: sessionID,
	}, nil
}

func (s *AuthService) VerifyCode(sessionID, code string) (string, error) {
	user, _ := s.UserRepository.GetUserBySessionID(sessionID)
	if user == nil {
		return "", errors.New(ErrUserNotFound)
	}
	if user.Code == code {
		claims := &jwt.Claims{
			UserID:    strconv.FormatUint(uint64(user.ID), 10),
			SessionID: sessionID,
			Phone:     user.Phone,
		}
		token, err := jwt.NewJWT(os.Getenv("JWT_SECRET")).GenerateToken(claims)
		if err != nil {
			return "", err
		}
		return token, nil
	}
	return "", errors.New(ErrInvalidCredentials)
}
