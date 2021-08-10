package comments

import "testing"

type testpair struct {
	value string
	res   Comments
}

type testpair2 struct {
	commentID string
	value     map[string]interface{}
	res       Comments
}

var testsReturnComment = []testpair{
	{"1", Comments{ID: 1, PostID: 1, Name: "jinzu", Email: "jinzu@gmail.com", Body: "zzzzzz"}},
	{"2", Comments{ID: 2, PostID: 1, Name: "jinzu", Email: "jinzu@gmail.com", Body: "zzzzzz"}},
	{"3", Comments{ID: 3, PostID: 1, Name: "jinzu", Email: "jinzu@gmail.com", Body: "zzzzzz"}},
}

var testsCreateComment = []testpair2{
	{"4", map[string]interface{}{"ID": 0, "PostID": 1, "Name": "jinzu", "Email": "jinzu@gmail.com", "Body": "zzzzzz"}, Comments{ID: 4, PostID: 1, Name: "jinzu", Email: "jinzu@gmail.com", Body: "zzzzzz"}},
}

var testsUpdateComment = []testpair2{
	{"4", map[string]interface{}{"ID": 4, "PostID": 1, "Name": "jinzu", "Email": "jinzu@gmail.com", "Body": "hi"}, Comments{ID: 4, PostID: 1, Name: "jinzu", Email: "jinzu@gmail.com", Body: "hi"}},
}

func TestGetAllComments(t *testing.T) {
	v := GetAllComments()
	for i := 0; i < 3; i++ {
		if v[i] != testsReturnComment[i].res {
			t.Error("Error")
		}
	}
}

func TestGetComment(t *testing.T) {
	for _, pair := range testsReturnComment {
		v, _ := GetComment(pair.value)
		if v != pair.res {
			t.Error(
				"For", pair.value,
				"expected", pair.res,
				"got", v,
			)
		}
	}
}

func TestCreateComment(t *testing.T) {
	CreateComment(testsCreateComment[0].value)
	v, _ := GetComment(testsCreateComment[0].commentID)
	if v != testsCreateComment[0].res {
		t.Error("Error",
			"expected", testsCreateComment[0].res,
			"got", v,
		)
	}
}

func TestUpdateComment(t *testing.T) {
	UpdateComment(testsUpdateComment[0].value, testsUpdateComment[0].commentID)
	v, _ := GetComment(testsUpdateComment[0].commentID)
	if v != testsUpdateComment[0].res {
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
