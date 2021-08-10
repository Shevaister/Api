package users

import (
	"Api/pkg/db"
)

type Users struct {
	ID       int
	Email    string
	Password string
}

func init() {
	db.DB.AutoMigrate(&Users{})
}

// GetUser for user registration or/and authorization through facebook/google
func GetUser(email string) (user Users) {
	err := db.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		user.Email = email
		db.DB.Select("email").Create(&user)
	} else {
		return
	}
	db.DB.Where("email = ?", email).First(&user)
	return
}

// Login For user authorization with email and password
func Login(email string, password string) (err error) {
	var user Users
	err = db.DB.Where("email = ?", email).Where("password = ?", password).First(&user).Error
	return
}

// Register For user registration with email and password
func Register(email string, password string) bool {
	var user Users
	err := db.DB.Where("email = ?", email).First(&user).Error
	if err == nil {
		if user.Password != "" {
			return false
		}
		user.Password = password
		db.DB.Save(&user)
		return true
	}
	user.Email = email
	user.Password = password
	db.DB.Select("email", "password").Create(&user)
	return true
}
