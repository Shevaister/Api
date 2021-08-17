package comments

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var comment = []Comments{{ID: 1, PostID: 1, Name: "jinzu", Email: "jinzu@gmail.com", Body: "zzzzzz"}, {ID: 2, PostID: 1, Name: "jinzu", Email: "jinzu@gmail.com", Body: "zzzzzz"}, {ID: 3, PostID: 1, Name: "jinzu", Email: "jinzu@gmail.com", Body: "zzzzzz"}}

func TestGetAllComments(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	commentRepo := ProvideCommentRepostiory(gdb)

	mock.ExpectQuery(
		"SELECT * FROM `comments`").
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "post_id", "name", "email", "body"}).
				AddRow(comment[0].ID, comment[0].PostID, comment[0].Name, comment[0].Email, comment[0].Body).
				AddRow(comment[1].ID, comment[1].PostID, comment[1].Name, comment[1].Email, comment[1].Body).
				AddRow(comment[2].ID, comment[2].PostID, comment[2].Name, comment[2].Email, comment[2].Body))

	res := commentRepo.GetAllComments()

	require.Equal(t, comment, res)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetComment(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	commentRepo := ProvideCommentRepostiory(gdb)

	mock.ExpectQuery(
		"SELECT * FROM `comments` WHERE `comments`.`id` = ? ORDER BY `comments`.`id` LIMIT 1").
		WithArgs("3").
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "post_id", "name", "email", "body"}).
				AddRow(comment[2].ID, comment[2].PostID, comment[2].Name, comment[2].Email, comment[2].Body))

	res, err := commentRepo.GetComment("3")

	require.NoError(t, err)
	require.Equal(t, comment[2], res)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateComment(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	comment := map[string]interface{}{"id": 3, "PostID": 1, "name": "jinzu", "email": "jinzu@gmail.com", "body": "zzzzzz"}

	commentRepo := ProvideCommentRepostiory(gdb)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `comments` (`post_id`,`body`,`email`,`id`,`name`) VALUES (?,?,?,?,?)").WithArgs(comment["PostID"], comment["body"], comment["email"], comment["id"], comment["name"]).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	commentRepo.CreateComment(comment)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateComment(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	comment := map[string]interface{}{"id": 3, "PostID": 1, "name": "jinzyy", "email": "jinzu@gmail.com", "body": "zzzzzz"}

	commentRepo := ProvideCommentRepostiory(gdb)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `comments` SET `post_id`=?,`body`=?,`email`=?,`name`=? WHERE `id` = ?").WithArgs(comment["PostID"], comment["body"], comment["email"], comment["name"], comment["id"]).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	commentRepo.UpdateComment(comment)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteComment(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	commentRepo := ProvideCommentRepostiory(gdb)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM `comments` WHERE `comments`.`id` = ?").WithArgs("3").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	commentRepo.DeleteComment("3")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
