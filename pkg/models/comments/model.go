package comments

type Comments struct {
	ID     int
	PostID uint `json:"PostID"`
	Name   string
	Email  string
	Body   string
}
