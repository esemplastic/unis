package unis

import (
	"testing"
)

func TestJoinerFunc_Join(t *testing.T) {
	type args struct {
		part1 string
		part2 string
	}
	tests := []struct {
		name string
		j    JoinerFunc
		args args
		want string
	}{
		{
			"join1",
			NewJoiner("/"),
			args{
				"file",
				"path",
			},
			"file/path",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.Join(tt.args.part1, tt.args.part2); got != tt.want {
				t.Errorf("JoinerFunc.Join() = %v, want %v", got, tt.want)
			}
		})
	}
}
