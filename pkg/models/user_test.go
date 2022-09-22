package models

import "testing"

func TestUsersGetById(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	repo := &UserModel{Client: client}

	results, err := repo.Get("632655b353adec83f7f2d6a5")

	if err != nil {
		t.Fatalf(err.Error())
	}

	if results == nil {
		t.Fatalf("User should have been returned")
	}

	expectedUsername := "test"
	if results.Username != expectedUsername {
		t.Fatalf("Expected username does not match")
	}
}
