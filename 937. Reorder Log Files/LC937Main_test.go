package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reorderLogFiles(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		logs []string
	}
	tests := []struct {
		name    string
		args    args
		expects []string
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				logs: []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"},
			},
			expects: []string{"g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"},
		},
		{
			name: "second",
			args: args{
				logs: []string{"1 n u", "r 527", "j 893", "6 14", "6 82"},
			},
			expects: []string{"1 n u", "r 527", "j 893", "6 14", "6 82"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(tt.expects, reorderLogFiles(tt.args.logs))
		})
	}
}
