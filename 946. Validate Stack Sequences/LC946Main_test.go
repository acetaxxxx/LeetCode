package LC946

import (
	"testing"
)

func Test_validateStackSequences(t *testing.T) {
	type args struct {
		pushed []int
		popped []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateStackSequences(tt.args.pushed, tt.args.popped); got != tt.want {
				t.Errorf("validateStackSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indexOf(t *testing.T) {
	type args struct {
		a []int
		t int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases
		{
			name: "success",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				t: 3,
			},
			want: 2,
		},
		{
			name: "cant find",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				t: 6,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := indexOf(tt.args.a, tt.args.t); got != tt.want {
				t.Errorf("indexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
