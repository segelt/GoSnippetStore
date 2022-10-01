package models

import (
	"testing"
)

func TestSnippetsByUserFilter(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

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

	repo := &SnippetModel{Client: client}
	for _, tc := range tests {
		filter := SnippetFilter{
			UserId: &tc.userId,
		}
		res, err := repo.ByUser(filter)
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

func TestSnippetsRetrievedById(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	type test struct {
		snippetId      string
		resultReturned bool
		errReturned    bool
		name           string
	}

	tests := []test{
		{snippetId: "63268461ea977bd27abcfcd7", resultReturned: true, errReturned: false, name: "Valid snippet via ID"},
		{snippetId: "111111111111111111111111", resultReturned: false, errReturned: true, name: "Invalid snippet ID"},
	}

	repo := &SnippetModel{Client: client}
	for _, tc := range tests {
		res, err := repo.Single(tc.snippetId)
		t.Logf("At test case %s", tc.name)
		if (err != nil) != tc.errReturned {
			t.Fatalf("Expected errReturned: %t. Got: %t", tc.errReturned, (err != nil))
		}

		if (res != nil) != tc.resultReturned {
			t.Fatalf("Expected resultReturned: %t. Got %t", tc.resultReturned, (res != nil))
		}
	}
}
