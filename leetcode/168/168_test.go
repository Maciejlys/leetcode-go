package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		column int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{
				column: 1,
			},
			want: "A",
		},
		{
			name: "example 2",
			args: args{
				column: 28,
			},
			want: "AB",
		},
		{
			name: "example 3",
			args: args{
				column: 701,
			},
			want: "ZY",
		},
		{
			name: "example 4",
			args: args{
				column: 2147483647,
			},
			want: "FXSHRXW",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := convertToTitle(tt.args.column); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
