package lc473

import "testing"

func Test_makesquare(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				nums: []int{4, 5, 6, 3, 2},
			},
			want: false,
		},
		{
			name: "second",
			args: args{
				nums: []int{1,1,2,2,2},
			},
			want: true,
		},
		{
			name: "third",
			args: args{
				nums: []int{3,3,3,3,4},
			},
			want: false,
		},
		{
			name: "fourth",
			args: args{
				nums: []int{},
			},
			want: false,
		},
		{
			name: "fifth",
			args: args{
				nums: []int{10,6,5,5,5,3,3,3,2,2,2,2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makesquare(tt.args.nums); got != tt.want {
				t.Errorf("makesquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
