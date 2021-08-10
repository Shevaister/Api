package posts

import (
	"Api/pkg/db"
)

type Posts struct {
	ID     int
	UserID uint `json:"UserID"`
	Title  string
	Body   string
}

func init() {
	db.DB.AutoMigrate(&Posts{})
}

func GetAllPosts() (post []Posts) {
	db.DB.Find(&post)
	return post
}

func GetPost(id string) (post Posts, err error) {
	err = db.DB.First(&post, id).Error
	return
}

func CreatePost(reqBody map[string]interface{}) {
	db.DB.Model(&Posts{}).Create(&reqBody)
}

func UpdatePost(reqBody map[string]interface{}, id string) (err error) {
	var post Posts
	err = db.DB.First(&post, id).Error
	if err != nil {
		return
	}
	db.DB.Model(&post).Updates(reqBody)
	return
}

func DeletePost(id string) {
	db.DB.Delete(&Posts{}, id)
}
