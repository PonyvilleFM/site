package mockdb

import (
	"testing"

	"github.com/PonyvilleFM/site/db"
)

func TestUserDao(t *testing.T) {
	ud := NewUserDao()

	_, err := ud.Register(db.User{
		Username:      "AzureDiamond",
		Email:         "its@only.stars.to.me",
		IsAdmin:       true,
		IsDJ:          true,
		TwitterHandle: "@AzureDiamond",
	}, "hunter2")
	if err != nil {
		t.Fatalf("expected ud.Register to return no error, got: %v", err)
	}

	ok, _, err := ud.Login("AzureDiamond", "hunter2")
	if err != nil {
		t.Fatalf("ud.Login failed: %v", err)
	}
	if !ok {
		t.Fatalf("expected ud.Login to work, it didn't")
	}

	_, _, err = ud.Login("Whip", "Nae Nae")
	if err != db.ErrNoSuchUser {
		t.Fatalf("invalid user was able to log in")
	}

	_, err = ud.Info("AzureDiamond")
	if err != nil {
		t.Fatalf("expected ud.Info to work, it didn't")
	}

	_, err = ud.Info("Whip")
	if err != db.ErrNoSuchUser {
		t.Fatalf("invalid user information was able to be fetched")
	}
}
