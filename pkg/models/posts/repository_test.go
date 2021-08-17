package posts

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/stretchr/testify/require"
)

var post = []Posts{{ID: 5, UserID: 1, Title: "sda", Body: "dsa"}, {ID: 6, UserID: 1, Title: "smt", Body: "dsa"}}

func TestGetAllPosts(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	postRepo := ProvidePostRepostiory(gdb)

	mock.ExpectQuery(
		"SELECT * FROM `posts`").
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_id", "title", "body"}).
				AddRow(post[0].ID, post[0].UserID, post[0].Title, post[0].Body).
				AddRow(post[1].ID, post[1].UserID, post[1].Title, post[1].Body))

	res := postRepo.GetAllPosts()

	require.Equal(t, post, res)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetPost(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	postRepo := ProvidePostRepostiory(gdb)

	mock.ExpectQuery(
		"SELECT * FROM `posts` WHERE `posts`.`id` = ? ORDER BY `posts`.`id` LIMIT 1").
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_id", "title", "body"}).
				AddRow(post[0].ID, post[0].UserID, post[0].Title, post[0].Body))

	res, err := postRepo.GetPost("5")

	require.NoError(t, err)
	require.Equal(t, post[0], res)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreatePost(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	post := map[string]interface{}{"id": 1, "UserID": 1, "title": "hg", "body": "tu"}

	postRepo := ProvidePostRepostiory(gdb)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `posts` (`user_id`,`body`,`id`,`title`) VALUES (?,?,?,?)").WithArgs(post["UserID"], post["body"], post["id"], post["title"]).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	postRepo.CreatePost(post)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdatePost(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	post := map[string]interface{}{"id": 1, "UserID": 1, "title": "hg", "body": "tu"}

	postRepo := ProvidePostRepostiory(gdb)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `posts` SET `user_id`=?,`body`=?,`title`=? WHERE `id` = ?").WithArgs(post["UserID"], post["body"], post["title"], post["id"]).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	postRepo.UpdatePost(post)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeletePost(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	postRepo := ProvidePostRepostiory(gdb)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM `posts` WHERE `posts`.`id` = ?").WithArgs("3").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	postRepo.DeletePost("3")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
