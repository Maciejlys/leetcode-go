package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		str1 string
		str2 string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{
				str1: "ABCABC", str2: "ABC",
			},
			want: "ABC",
		},
		{
			name: "example 2",
			args: args{
				str1: "ABABAB", str2: "ABAB",
			},
			want: "AB",
		},
		{
			name: "example 3",
			args: args{
				str1: "TAUXXTAUXXTAUXXTAUXXTAUXX", str2: "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX",
			},
			want: "TAUXX",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := gcdOfStrings(tt.args.str1, tt.args.str2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
