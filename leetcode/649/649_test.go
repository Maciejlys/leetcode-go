package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		senate string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{
				"RD",
			},
			want: "Radiant",
		},
		{
			name: "example 2",
			args: args{
				"RDD",
			},
			want: "Dire",
		},
		{
			name: "example 3",
			args: args{
				"DDRRRR",
			},
			want: "Radiant",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := predictPartyVictory(tt.args.senate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
