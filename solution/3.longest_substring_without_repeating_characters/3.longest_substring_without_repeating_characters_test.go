/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

package solution

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{s: "bbbbbb"},
			want: 1,
		},
		{
			args: args{s: "abcabcbb"},
			want: 3,
		},
		{
			args: args{s: ""},
			want: 0,
		},
		{
			args: args{s: "pwwlew"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
