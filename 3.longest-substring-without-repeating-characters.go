package main

import (
	"strings"
)

// @leet start
func lengthOfLongestSubstring(s string) int {
	longest := 0
	buff := ""
	for _, r := range s {
		idx := strings.IndexRune(buff, r)
		buff += string(r)
		if idx >= 0 {
			buff = buff[idx+1:]
		}
		if l := len(buff); l > longest {
			longest = l
		}
	}
	return longest
}

// @leet end

