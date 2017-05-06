package unis

import (
	"strings"
	"testing"
)

type staticPathResolver struct {
	paramStart    byte
	wildcardStart byte
}

func newStaticPathResolver(paramStartSymbol, wildcardStartParamSymbol byte) Processor {
	return staticPathResolver{
		paramStart:    paramStartSymbol,
		wildcardStart: wildcardStartParamSymbol,
	}
}

func (s staticPathResolver) Process(original string) (result string) {
	i := strings.IndexByte(original, s.paramStart)
	v := strings.IndexByte(original, s.wildcardStart)

	return NewConditional(NewRangeEnd(i),
		NewRangeEnd(v)).Process(original)
}

var resolveStaticPath = newStaticPathResolver(':', '*')

func TestConditional(t *testing.T) {
	tests := []struct {
		original string
		result   string
	}{
		{"/api/users/:id", "/api/users/"},
		{"/public/assets/*file", "/public/assets/"},
		{"/profile/:id/files/*file", "/profile/"},
	}

	for i, tt := range tests {
		if expected, got := tt.result, resolveStaticPath.Process(tt.original); expected != got {
			t.Fatalf("[%d] - expected '%s' but got '%s'", i, expected, got)
		}
	}
}
