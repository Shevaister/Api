package users

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func ProvideUserRepostiory(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

// GetUser for user registration or/and authorization through facebook/google
func (u *UserRepository) GetUser(email string) (user Users) {
	err := u.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		user.Email = email
		u.DB.Select("email").Create(&user)
	} else {
		return
	}
	return
}

// Login For user authorization with email and password
func (u *UserRepository) Login(email string, password string) (err error) {
	var user Users
	err = u.DB.Where("email = ?", email).Where("password = ?", password).First(&user).Error
	return
}

// Register For user registration with email and password
func (u *UserRepository) Register(email string, password string) bool {
	var user Users
	err := u.DB.Where("email = ?", email).First(&user).Error
	if err == nil {
		if user.Password != "" {
			return false
		}
		user.Password = password
		u.DB.Save(&user)
		return true
	}
	user.Email = email
	user.Password = password
	u.DB.Select("email", "password").Create(&user)
	return true
}
