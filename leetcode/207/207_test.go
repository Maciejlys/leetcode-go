package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		numCourses int
		prer       [][]int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				numCourses: 2,
				prer:       [][]int{{1, 0}},
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				numCourses: 2,
				prer:       [][]int{{1, 0}, {0, 1}},
			},
			want: false,
		},
		{
			name: "example 3",
			args: args{
				numCourses: 2,
				prer:       [][]int{{0, 1}},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := canFinish(tt.args.numCourses, tt.args.prer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
