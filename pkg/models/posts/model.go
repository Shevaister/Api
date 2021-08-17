package posts

type Posts struct {
	ID     int
	UserID uint `json:"UserID"`
	Title  string
	Body   string
}
