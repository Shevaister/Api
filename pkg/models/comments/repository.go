package comments

import (
	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func ProvideCommentRepostiory(DB *gorm.DB) CommentRepository {
	return CommentRepository{DB: DB}
}

func (c *CommentRepository) GetAllComments() (comment []Comments) {
	c.DB.Find(&comment)
	return
}

func (c *CommentRepository) GetComment(id string) (comment Comments, err error) {
	err = c.DB.First(&comment, id).Error
	return
}

func (c *CommentRepository) CreateComment(reqBody map[string]interface{}) {
	c.DB.Model(&Comments{}).Create(reqBody)
}

func (c *CommentRepository) UpdateComment(reqBody map[string]interface{}) (err error) {
	var comment Comments
	comment.ID = reqBody["id"].(int)
	delete(reqBody, "id")
	err = c.DB.Model(&comment).Updates(reqBody).Error
	return
}

func (c *CommentRepository) DeleteComment(id string) {
	c.DB.Delete(&Comments{}, id)
}
