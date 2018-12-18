package main

import (
	"fmt"
)

/**
	给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
	给定 "abcabcbb" ，没有重复字符的最长子串是 "abc" ，那么长度就是3。
	给定 "bbbbb" ，最长的子串就是 "b" ，长度是1。
	给定 "pwwkew" ，最长子串是 "wke" ，长度是3。请注意答案必须是一个子串，"pwke" 是 子序列 而不是子串。
 */

func main() {
	ret := lengthOfLongestSubstring("aabcdabcde")
	fmt.Println("**", ret, "**")
}

func lengthOfLongestSubstring(s string) int {
	location := [256]int{} 	//只有256长是因为，假定输入的字符串只有ASCII字符
	for i := range location {
		location[i] = -1 	//先设置所有的字符都没有见过
	}

	maxLen, left := 0, 0
	for i := 0; i < len(s); i++ {
		//fmt.Println(s[i], "location：", location[s[i]], "left：", left)
		if location[s[i]] >= left {
			left = location[s[i]] + 1 // 在s[left:i+1]中去除s[i]字符及其之前的部分
			fmt.Println(s[i], "location：", location[s[i]], "left：", left)
		} else if i+1-left > maxLen {
			//fmt.Println(s[left:i+1])
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}

	return maxLen
}
