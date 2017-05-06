package unis

import (
	"path"
	"strings"
	"testing"
)

func newPathNormalizer() Processor {
	slash := "/"
	replacer := NewReplacer(map[string]string{
		`\`:  slash,
		`//`: slash,
	})

	suffixRemover := NewSuffixRemover(slash)
	slashPrepender := NewTargetedJoiner(0, slash[0])

	toLower := ProcessorFunc(strings.ToLower)
	pathCleaner := ProcessorFunc(path.Clean)
	return NewChain(
		replacer,
		suffixRemover,
		slashPrepender,
		pathCleaner,
		toLower,
	)
}

var defaultPathNormalizer = newPathNormalizer()

func normalizePath(path string) string {
	if path == "" {
		return path
	}
	return defaultPathNormalizer.Process(path)
}

func TestChain(t *testing.T) {
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
