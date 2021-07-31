package models

import (
	"fmt"
	"strconv"
	"testing"
)

type testpair struct {
	value string
	res   map[string]interface{}
}

var testsReturnComment = []testpair{
	{"1", map[string]interface{}{"ID": 1, "PostID": 1, "Name": "jinzu", "Email": "jinzu@gmail.com", "Body": "zzzzzz"}},
	{"2", map[string]interface{}{"ID": 2, "PostID": 1, "Name": "jinzu", "Email": "jinzu@gmail.com", "Body": "zzzzzz"}},
	{"3", map[string]interface{}{"ID": 3, "PostID": 1, "Name": "jinzu", "Email": "jinzu@gmail.com", "Body": "zzzzzz"}},
}

var testsReturnPost = []testpair{
	{"1", map[string]interface{}{"ID": 1, "UserID": 1, "Title": "not", "Body": "dsadasdadwaa"}},
	{"2", map[string]interface{}{"ID": 2, "UserID": 1, "Title": "not", "Body": "dsadasdadwaa"}},
}

var testsGetUser = []testpair{
	{"djio@gmail.com", map[string]interface{}{"ID": 1, "Email": "djio@gmail.com"}},
	{"jinzu@gmail.com", map[string]interface{}{"ID": 2, "Email": "jinzu@gmail.com"}},
}

var testsCreateComment = []testpair{
	{"4", map[string]interface{}{"ID": 0, "PostID": 1, "Name": "jinzu", "Email": "jinzu@gmail.com", "Body": "zzzzzz"}},
}

var testsCreatePost = []testpair{
	{"3", map[string]interface{}{"ID": 0, "UserID": 1, "Title": "not", "Body": "dsadasdadwaa"}},
}

var testsUpdateComment = []testpair{
	{"4", map[string]interface{}{"ID": 4, "PostID": 1, "Name": "jinzu", "Email": "jinzu@gmail.com", "Body": "hi"}},
}

var testsUpdatePost = []testpair{
	{"3", map[string]interface{}{"ID": 3, "UserID": 1, "Title": "not", "Body": "hello"}},
}

var testsRegisterLogin = []testpair{
	{"3", map[string]interface{}{"Email": "smt@gmail.com", "Password": "sssy"}},
}

func TestGetAllComments(t *testing.T) {
	v := GetAllComments()
	for i := 0; i < 3; i++ {
		if int(v[i].ID) != testsReturnComment[i].res["ID"] || int(v[i].PostID) != testsReturnComment[i].res["PostID"] || v[i].Name != testsReturnComment[i].res["Name"] || v[i].Email != testsReturnComment[i].res["Email"] || v[i].Body != testsReturnComment[i].res["Body"] {
			t.Error("Error")
		}
	}
}

func TestGetAllPosts(t *testing.T) {
	v := GetAllPosts()
	for i := 0; i < 2; i++ {
		if int(v[i].ID) != testsReturnPost[i].res["ID"] || int(v[i].UserID) != testsReturnPost[i].res["UserID"] || v[i].Title != testsReturnPost[i].res["Title"] || v[i].Body != testsReturnPost[i].res["Body"] {
			t.Error("Error")
		}
	}
}

func TestGetComment(t *testing.T) {
	for _, pair := range testsReturnComment {
		v, _ := GetComment(pair.value)
		if int(v.ID) != pair.res["ID"] || int(v.PostID) != pair.res["PostID"] || v.Name != pair.res["Name"] || v.Email != pair.res["Email"] || v.Body != pair.res["Body"] {
			t.Error(
				"For", pair.value,
				"expected", pair.res,
				"got", v,
			)
		}
	}
}

func TestGetPost(t *testing.T) {
	for _, pair := range testsReturnPost {
		v, _ := GetPost(pair.value)
		if int(v.ID) != pair.res["ID"] || int(v.UserID) != pair.res["UserID"] || v.Title != pair.res["Title"] || v.Body != pair.res["Body"] {
			t.Error(
				"For", pair.value,
				"expected", pair.res,
				"got", v,
			)
		}
	}
}

func TestGetUser(t *testing.T) {
	for _, pair := range testsGetUser {
		v := GetUser(pair.value)
		if int(v.ID) != pair.res["ID"] || v.Email != pair.res["Email"] {
			fmt.Println(v.ID)
			fmt.Println(pair.res["ID"])
			t.Error(
				"For", pair.value,
				"expected", pair.res,
				"got", v,
			)
		}
	}
}

func TestCreateComment(t *testing.T) {
	CreateComment(testsCreateComment[0].res)
	v, _ := GetComment(testsCreateComment[0].value)
	if strconv.Itoa(int(v.ID)) != testsCreateComment[0].value || int(v.PostID) != testsCreateComment[0].res["PostID"] || v.Name != testsCreateComment[0].res["Name"] || v.Email != testsCreateComment[0].res["Email"] || v.Body != testsCreateComment[0].res["Body"] {
		t.Error("Error",
			"expected", testsCreateComment[0].res,
			"got", v,
		)
	}
}

func TestCreatePost(t *testing.T) {
	CreatePost(testsCreatePost[0].res)
	v, _ := GetPost(testsCreatePost[0].value)
	if strconv.Itoa(int(v.ID)) != testsCreatePost[0].value || int(v.UserID) != testsCreatePost[0].res["UserID"] || v.Title != testsCreatePost[0].res["Title"] || v.Body != testsCreatePost[0].res["Body"] {
		t.Error("Error",
			"expected", testsCreatePost[0].res,
			"got", v,
		)
	}

}

func TestUpdateComment(t *testing.T) {
	UpdateComment(testsUpdateComment[0].res, testsUpdateComment[0].value)
	v, _ := GetComment(testsUpdateComment[0].value)
	if int(v.ID) != testsUpdateComment[0].res["ID"] || int(v.PostID) != testsUpdateComment[0].res["PostID"] || v.Name != testsUpdateComment[0].res["Name"] || v.Email != testsUpdateComment[0].res["Email"] || v.Body != testsUpdateComment[0].res["Body"] {
		t.Error("Error")
	}

}

func TestUpdatePost(t *testing.T) {
	UpdatePost(testsUpdatePost[0].res, testsUpdatePost[0].value)
	v, _ := GetPost(testsUpdatePost[0].value)
	if int(v.ID) != testsUpdatePost[0].res["ID"] || int(v.UserID) != testsUpdatePost[0].res["UserID"] || v.Title != testsUpdatePost[0].res["Title"] || v.Body != testsUpdatePost[0].res["Body"] {
		t.Error("Error")
	}
}

func TestDeleteComment(t *testing.T) {
	DeleteComment("4")
	_, err := GetComment("4")
	if err == nil {
		t.Error("Error")
	}
}

func TestDeletePost(t *testing.T) {
	DeletePost("3")
	_, err := GetPost("3")
	if err == nil {
		t.Error("Error")
	}
}

func TestRegisterLogin(t *testing.T) {
	Register(testsRegisterLogin[0].res["Email"].(string), testsRegisterLogin[0].res["Password"].(string))
	err := Login(testsRegisterLogin[0].res["Email"].(string), testsRegisterLogin[0].res["Password"].(string))
	if err != nil {
		t.Error("Error")
	}
}
