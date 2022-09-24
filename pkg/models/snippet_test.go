package models

import (
	"testing"
)

func TestSnippetsByUserFilter(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	repo := &SnippetModel{Client: client}

	type test struct {
		userId      string
		resultCount int
		errReturned bool
		name        string
	}

	tests := []test{
		{userId: "632655b353adec83f7f2d6a5", resultCount: 7, errReturned: false, name: "Valid user with multiple snippets test"},
		{userId: "632655b353adec83f7f2d6a51", resultCount: 0, errReturned: false, name: "User does not exist"},
	}

	for _, tc := range tests {
		res, err := repo.ByUser(tc.userId)
		t.Logf("At test case %s", tc.name)
		if (err != nil) != tc.errReturned {
			t.Fatalf("Expected errReturned: %t. Got: %t", tc.errReturned, (err != nil))
		}

		if res != nil && len(res) != tc.resultCount {
			t.Fatalf("Expected returned snippet count and the actual count does not match. Expected count: %d, got: %d", tc.resultCount, len(res))
		}

		if res == nil && tc.resultCount > 0 {
			t.Fatalf("Expected count: %d, got nil result.", tc.resultCount)
		}
	}

}
