package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		strs []string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{
				strs: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
		{
			name: "example 2",
			args: args{
				strs: []string{"dog", "racecar", "car"},
			},
			want: "",
		},
		{
			name: "example 3",
			args: args{
				strs: []string{""},
			},
			want: "",
		},
		{
			name: "example 4",
			args: args{
				strs: []string{"a"},
			},
			want: "a",
		},
		{
			name: "example 5",
			args: args{
				strs: []string{"", "b"},
			},
			want: "",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := longestCommonPrefix(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
