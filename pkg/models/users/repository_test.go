package users

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	rs := sqlmock.NewRows([]string{"id", "email", "password"})
	//zs := rs.AddRow(1, "deg@gmail.com", "")

	commentRepo := ProvideUserRepostiory(gdb)

	mock.ExpectQuery("SELECT * FROM `users` WHERE email = ? ORDER BY `users`.`id` LIMIT 1").WithArgs("deg@gmail.com").WillReturnRows(rs)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users` (`email`) VALUES (?)").WithArgs("deg@gmail.com").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	commentRepo.GetUser("deg@gmail.com")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestLogin(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	rs := sqlmock.NewRows([]string{"id", "email", "password"})

	commentRepo := ProvideUserRepostiory(gdb)

	mock.ExpectQuery("SELECT * FROM `users` WHERE email = ? AND (password = ?) ORDER BY `users`.`id` LIMIT 1").WithArgs("deg@gmail.com", "smt").WillReturnRows(rs)

	commentRepo.Login("deg@gmail.com", "smt")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRegister(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	rs := sqlmock.NewRows([]string{"id", "email", "password"})

	commentRepo := ProvideUserRepostiory(gdb)

	mock.ExpectQuery("SELECT * FROM `users` WHERE email = ? ORDER BY `users`.`id` LIMIT 1").WithArgs("deg@gmail.com").WillReturnRows(rs)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users` (`email`,`password`) VALUES (?,?)").WithArgs("deg@gmail.com", "smt").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	commentRepo.Register("deg@gmail.com", "smt")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
