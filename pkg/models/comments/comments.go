package comments

import (
	"Api/pkg/db"
)

type Comments struct {
	ID     int
	PostID uint `json:"PostID"`
	Name   string
	Email  string
	Body   string
}

func init() {
	db.DB.AutoMigrate(&Comments{})
}

func GetAllComments() (comment []Comments) {
	db.DB.Find(&comment)
	return
}

func GetComment(id string) (comment Comments, err error) {
	err = db.DB.First(&comment, id).Error
	return
}

func CreateComment(reqBody map[string]interface{}) {
	db.DB.Model(&Comments{}).Create(reqBody)
}

func UpdateComment(reqBody map[string]interface{}, id string) (err error) {
	var comment Comments
	err = db.DB.First(&comment, id).Error
	if err != nil {
		return
	}
	db.DB.Model(&comment).Updates(reqBody)
	return
}

func DeleteComment(id string) {
	db.DB.Delete(&Comments{}, id)
}
