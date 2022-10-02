package delivery

import (
	"testing"
)

func TestUserProto(t *testing.T) {
	user := User{
		Username: "godpeny",
	}
	if user.GetUsername() != "godpeny" {
		t.Fatal("expected user name to be godpeny")
	}
}
