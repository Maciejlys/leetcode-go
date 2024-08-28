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
		want [][]string
	}{
		{
			name: "example 1",
			args: args{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{{"tan", "nat"}, {"bat"}, {"eat", "tea", "ate"}},
		},
		{
			name: "example 2",
			args: args{
				strs: []string{""},
			},
			want: [][]string{{""}},
		},
		{
			name: "example 3",
			args: args{
				strs: []string{"a"},
			},
			want: [][]string{{"a"}},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
