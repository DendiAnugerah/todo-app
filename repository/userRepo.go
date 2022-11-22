package repository

import (
	"github.com/DendiAnugerah/Todo/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) AddUser(user model.User) error {
	return u.db.Create(&user).Error
}

func (u *UserRepository) CheckPasswordLength(pass string) bool {
	if len(pass) <= 5 {
		return false
	}
	return true
}

func (u *UserRepository) IsUsernameAvail(user model.User) error {
	return u.db.Where("username = ? ", user.Username).First(&user).Error
}
