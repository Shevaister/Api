package users

import "testing"

type testpair struct {
	value string
	res   Users
}

type testpair2 struct {
	userID string
	value  map[string]interface{}
	res    Users
}

var testsGetUser = []testpair{
	{"djio@gmail.com", Users{ID: 1, Email: "djio@gmail.com"}},
	{"jinzu@gmail.com", Users{ID: 2, Email: "jinzu@gmail.com"}},
}

var testsRegisterLogin = []testpair2{
	{"3", map[string]interface{}{"Email": "smt@gmail.com", "Password": "sssy"}, Users{Email: "smt@gmail.com", Password: "sssy"}},
}

func TestGetUser(t *testing.T) {
	for _, pair := range testsGetUser {
		v := GetUser(pair.value)
		if v != pair.res {
			t.Error(
				"For", pair.value,
				"expected", pair.res,
				"got", v,
			)
		}
	}
}

func TestRegisterLogin(t *testing.T) {
	Register(testsRegisterLogin[0].value["Email"].(string), testsRegisterLogin[0].value["Password"].(string))
	err := Login(testsRegisterLogin[0].value["Email"].(string), testsRegisterLogin[0].value["Password"].(string))
	if err != nil {
		t.Error("Error")
	}
}
