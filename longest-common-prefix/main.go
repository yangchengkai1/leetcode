package main

import "fmt"

func main() {
	s := []string{"flower", "flow", "flight"}

	fmt.Println(longestCommonPrefix(s))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	minLen := len(strs[0])
	index := 0

	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
			index = i
		}
	}

	for j := 0; j < len(strs[index]); j++ {
		for k := 0; k < len(strs); k++ {
			if strs[index][j] != strs[k][j] {
				return strs[index][:j]
			}
		}
	}

	return strs[index]
}
