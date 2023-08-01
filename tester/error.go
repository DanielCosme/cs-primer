package tester

import "testing"

func AssertNilErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("\nError is not nil, got: %v", err)
	}
}
