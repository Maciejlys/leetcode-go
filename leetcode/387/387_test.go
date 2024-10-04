package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				s: "leetcode",
			},
			want: 0,
		},
		{
			name: "example 2",
			args: args{
				s: "loveleetcode",
			},
			want: 2,
		},
		{
			name: "example 3",
			args: args{
				s: "aabb",
			},
			want: -1,
		},
		{
			name: "example 4",
			args: args{
				s: "z",
			},
			want: 0,
		},
		{
			name: "example 5",
			args: args{
				s: "bd",
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := firstUniqChar(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
