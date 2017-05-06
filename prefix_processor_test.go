package unis

import (
	"testing"
)

func TestPrefixRemover(t *testing.T) {
	tests := []struct {
		original string
		result   string
	}{
		{"/api/users/42", "api/users/42"},
		{"//api/users/42/", "api/users/42/"},
		{"api/users/", "api/users/"},
	}

	prefixRemover := NewPrefixRemover("/")

	for i, tt := range tests {
		if expected, got := tt.result, prefixRemover(tt.original); expected != got {
			t.Fatalf("[%d] - expected '%s' but got '%s'", i, expected, got)
		}
	}
}

func TestPrepender(t *testing.T) {
	tests := []struct {
		original string
		result   string
	}{
		{"/api/users/42", "/api/users/42"},
		{"//api/users\\42", "//api/users\\42"},
		{"api\\////users/", "/api\\////users/"},
	}
	prepender := NewPrepender("/")
	for i, tt := range tests {
		if expected, got := tt.result, prepender(tt.original); expected != got {
			t.Fatalf("[%d] - expected '%s' but got '%s'", i, expected, got)
		}
	}
}

func TestExclusivePrepender(t *testing.T) {
	tests := []struct {
		original string
		result   string
	}{
		{"/api/users/42", "/api/users/42"},
		// the only difference from simple Prepender is that this ExclusivePrepender
		// will make sure that we have only one slash as a prefix.
		{"//api/users\\42", "/api/users\\42"},
		{"api\\////users/", "/api\\////users/"},
	}
	prepender := NewExclusivePrepender("/")
	for i, tt := range tests {
		if expected, got := tt.result, prepender(tt.original); expected != got {
			t.Fatalf("[%d] - expected '%s' but got '%s'", i, expected, got)
		}
	}
}
