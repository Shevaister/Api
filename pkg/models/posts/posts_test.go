package posts

import "testing"

type testpair struct {
	value string
	res   Posts
}

type testpair2 struct {
	postID string
	value  map[string]interface{}
	res    Posts
}

var testsReturnPost = []testpair{
	{"1", Posts{ID: 1, UserID: 1, Title: "not", Body: "dsadasdadwaa"}},
	{"2", Posts{ID: 2, UserID: 1, Title: "not", Body: "dsadasdadwaa"}},
}

var testsCreatePost = []testpair2{
	{"3", map[string]interface{}{"ID": 0, "UserID": 1, "Title": "not", "Body": "dsadasdadwaa"}, Posts{ID: 3, UserID: 1, Title: "not", Body: "dsadasdadwaa"}},
}

var testsUpdatePost = []testpair2{
	{"3", map[string]interface{}{"ID": 3, "UserID": 1, "Title": "not", "Body": "hello"}, Posts{ID: 3, UserID: 1, Title: "not", Body: "hello"}},
}

func TestGetAllPosts(t *testing.T) {
	v := GetAllPosts()
	for i := 0; i < 2; i++ {
		if v[i] != testsReturnPost[i].res {
			t.Error("Error")
		}
	}
}

func TestGetPost(t *testing.T) {
	for _, pair := range testsReturnPost {
		v, _ := GetPost(pair.value)
		if v != pair.res {
			t.Error(
				"For", pair.value,
				"expected", pair.res,
				"got", v,
			)
		}
	}
}

func TestCreatePost(t *testing.T) {
	CreatePost(testsCreatePost[0].value)
	v, _ := GetPost(testsCreatePost[0].postID)
	if v != testsCreatePost[0].res {
		t.Error("Error",
			"expected", testsCreatePost[0].res,
			"got", v,
		)
	}

}

func TestUpdatePost(t *testing.T) {
	UpdatePost(testsUpdatePost[0].value, testsUpdatePost[0].postID)
	v, _ := GetPost(testsUpdatePost[0].postID)
	if v != testsUpdatePost[0].res {
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
