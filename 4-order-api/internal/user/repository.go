package user

import (
	"go/hw/4-order-api/pkg/db"

	"gorm.io/gorm/clause"
)

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{Database: database}
}

func (r *UserRepository) CreateUser(user *User) (*User, error) {
	err := r.Database.DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByPhone(phone string) (*User, error) {
	user := &User{}
	err := r.Database.DB.First(user, "phone = ?", phone).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user *User) (*User, error) {
	err := r.Database.DB.Clauses(clause.Returning{}).Updates(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserBySessionID(sessionID string) (*User, error) {
	user := &User{}
	err := r.Database.DB.First(user, "session_id = ?", sessionID).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
