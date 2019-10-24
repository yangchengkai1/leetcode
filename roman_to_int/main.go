package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(romanToInt("MCMXCVI"))
}

func romanToInt(s string) int {
	map1 := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	map2 := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	var num []int
	var number int

	for index, val := range map2 {
		if strings.Contains(s, index) {
			showCount := strings.Count(s, index)
			for i := 0; i < showCount; i++ {
				num = append(num, val)
			}
			s = strings.Replace(s, index, "", showCount)
		}
	}

	for _, val := range s {
		if map1[string(val)] != 0 {
			num = append(num, map1[string(val)])
		}
	}

	for _, value := range num {
		number += value
	}
	return number
}
