package server

import (
	"testing"

	"github.com/nathaponb/robusta-gosrv/internal/repository/user"
)

type TestUser struct {
	Username string
	Password string
}

func TestLogin(t *testing.T) {

	// define valid user
	validUser := user.User{
		FirstName:   "nathapon",
		LastName:    "b",
		DateOfBirth: "2023-29-03",
		UserName:    "nathaponb01",
		Email:       "nathaponb01@dev.com",
		Password:    "12345",
	}

	// define table of both true and false tests for testing result
	var table = []struct {
		username string
		password string
		want     bool
	}{
		{username: "nathaponb01", password: "1234", want: false},
		{username: "nathaponb02", password: "12345", want: false},
		{username: "nathaponb01", password: "12345", want: true},
	}

	for _, row := range table {
		if row.username != validUser.UserName && row.password != validUser.Password {
			t.Errorf("unauthorized user for %s", row.username)
		}
	}

	// or potentially implements test with httptest

}

func TestRegister(t *testing.T) {

	// define some existing user
	var dbUsers = []TestUser{
		{Username: "nathapoonb01", Password: "12345"},
		{Username: "nathapoonb02", Password: "12345"},
	}

	// define test cases of register request
	var regUsers = []struct {
		username string
		password string
		success  bool
	}{
		{username: "nathaponb01", password: "123456", success: false},
		{username: "nathaponb02", password: "12345", success: false},
		{username: "nathaponb03", password: "123456", success: true},
	}

	// check if the table of test request match already existed user expects to returns false
	for _, row := range regUsers {

		userExists(t, dbUsers, row.username)
	}

	// otherwise returns true
}

func userExists(t *testing.T, dbUsers []TestUser, username string) {
	found := false

	for _, dbUser := range dbUsers {
		if dbUser.Username == username {
			found = true
			break
		}
	}

	if found {
		t.Errorf("username %s already existed", username)
	}
}
