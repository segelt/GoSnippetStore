package models

import "testing"

func TestUsersGetById(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	repo := &UserModel{Client: client}

	type test struct {
		input           string
		userShouldBeNil bool
		errShouldBeNil  bool
		name            string
	}

	tests := []test{
		{input: "632655b353adec83f7f2d6a5", userShouldBeNil: false, errShouldBeNil: true, name: "Valid user test"},
		{input: "111111111111111111111111", userShouldBeNil: true, errShouldBeNil: false, name: "Non existing user test"},
	}

	for _, tc := range tests {
		got_result, got_err := repo.Get(tc.input)

		if (got_result == nil) != tc.userShouldBeNil {
			t.Fatalf("Test: %s. Expected user result to be %v, got: %v", tc.name, tc.userShouldBeNil, got_result)
		}

		if (got_err == nil) != tc.errShouldBeNil {
			t.Fatalf("Test: %s. Expected user result to be %v, got: %v", tc.name, tc.errShouldBeNil, got_err)
		}
	}
}

func TestUsersFilter(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	type test struct {
		inputFilter    UserFilter
		resultCount    int
		errShouldBeNil bool
		name           string
	}

	var fullUsername string = "test"
	var partialUsername string = "te"
	var partialInvalidUsername string = "abc"
	tests := []test{
		{
			inputFilter: UserFilter{Username: &fullUsername}, resultCount: 1, errShouldBeNil: true, name: "Full username filter test",
		},
		{
			inputFilter: UserFilter{Username: &partialUsername}, resultCount: 1, errShouldBeNil: true, name: "Partial username filter test",
		},
		{
			inputFilter: UserFilter{Username: &partialInvalidUsername}, resultCount: 0, errShouldBeNil: true, name: "Invalid username filter test",
		},
		{
			inputFilter: UserFilter{}, resultCount: 1, errShouldBeNil: true, name: "Empty filter",
		},
	}

	repo := &UserModel{Client: client}
	for _, tc := range tests {
		t.Logf("Running test %s\n", tc.name)
		result, err := repo.Filter(tc.inputFilter)

		if (err == nil) != tc.errShouldBeNil {
			t.Fatalf("Test %s. Returned error value does not match the desired error value. Current error value: %v. Wanted error value: %v", tc.name, err, tc.errShouldBeNil)
		}

		if *result != nil && len(*result) != tc.resultCount {
			t.Fatalf("Test %s. Mismatch on user amount returned. Wanted filtered user amount: %d. Got: %d", tc.name, tc.resultCount, len(*result))
		}

		if *result == nil && tc.resultCount != 0 {
			t.Fatalf("Test %s. Mismatch on user amount returned. Wanted %d users returned. Got no users returned", tc.name, tc.resultCount)
		}
	}
}
