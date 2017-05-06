package unis

import (
	"testing"
)

func TestTargetedJoiner(t *testing.T) {
	tests := []struct {
		original string
		result   string
	}{
		{"/api/users/42", "/api/users/42"},
		{"//api/users\\42", "/api/users/42"},
		{"api\\////users/", "/api/users"},
	}

	for i, tt := range tests {
		if expected, got := tt.result, normalizePath(tt.original); expected != got {
			t.Fatalf("[%d] - expected '%s' but got '%s'", i, expected, got)
		}
	}
}
