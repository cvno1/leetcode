/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */
package solution

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{"32"},
			want: 32,
		},
		{
			args: args{" -32"},
			want: -32,
		},
		{
			args: args{"000032"},
			want: 32,
		},
		{
			args: args{"32 other words"},
			want: 32,
		},
		{
			args: args{"other words 32"},
			want: 0,
		},
		{
			args: args{"2147483647"},
			want: 2147483647,
		},
		{
			args: args{"2147483648"},
			want: 2147483647,
		},
		{
			args: args{"21474836470"},
			want: 2147483647,
		},
		{
			args: args{"-91283472332"},
			want: -2147483648,
		},
		{
			args: args{"  0000000000012345678"},
			want: 12345678,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
