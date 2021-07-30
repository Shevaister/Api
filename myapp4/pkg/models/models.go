package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Comments struct {
	ID     uint
	PostID uint `json:"PostID"`
	Name   string
	Email  string
	Body   string
}

type Posts struct {
	ID     uint
	UserID uint `json:"UserID"`
	Title  string
	Body   string
}

type Users struct {
	ID       uint
	Email    string
	Password string
}

func init() {
	dsn := "mysql:mysql@tcp(127.0.0.1:3306)/parse?charset=utf8&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Posts{}, &Comments{}, &Users{})
}

//For comments

func GetAllComments() (comment []Comments) {
	db.Find(&comment)
	return
}

func GetComment(id string) (comment Comments, err error) {
	err = db.First(&comment, id).Error
	return
}

func CreateComment(reqBody map[string]interface{}) {
	db.Model(&Comments{}).Create(reqBody)
}

func UpdateComment(reqBody map[string]interface{}, id string) (err error) {
	var comment Comments
	err = db.First(&comment, id).Error
	if err != nil {
		return
	}
	db.Model(&comment).Updates(reqBody)
	return
}

func DeleteComment(id string) {
	db.Delete(&Comments{}, id)
}

//For posts

func GetAllPosts() (post []Posts) {
	db.Find(&post)
	return post
}

func GetPost(id string) (post Posts, err error) {
	err = db.First(&post, id).Error
	return
}

func CreatePost(reqBody map[string]interface{}) {
	db.Model(&Posts{}).Create(&reqBody)
}

func UpdatePost(reqBody map[string]interface{}, id string) (err error) {
	var post Posts
	err = db.First(&post, id).Error
	if err != nil {
		return
	}
	db.Model(&post).Updates(reqBody)
	return
}

func DeletePost(id string) {
	db.Delete(&Posts{}, id)
}

func GetUser(email string) (user Users) {
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		user.Email = email
		db.Select("email").Create(&user)
	} else {
		return
	}
	db.Where("email = ?", email).First(&user)
	return
}

func Login(email string, password string) (err error) {
	var user Users
	err = db.Where("email = ?", email).Where("password = ?", password).First(&user).Error
	return
}

func Register(email string, password string) bool {
	var user Users
	err := db.Where("email = ?", email).First(&user).Error
	if err == nil {
		if user.Password != "" {
			return false
		}
		user.Password = password
		db.Save(&user)
		return true
	}
	user.Email = email
	user.Password = password
	db.Select("email", "password").Create(&user)
	return true
}
