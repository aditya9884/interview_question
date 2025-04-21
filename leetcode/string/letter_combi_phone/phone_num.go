/*
Given a string containing digits from 2-9 inclusive, return all possible letter
combinations that the number could represent. Return the answer in any order.
A mapping of digits to letters (just like on the telephone buttons) is given below.
Note that 1 does not map to any letters.
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
Example 2:
Input: digits = ""
Output: []
Input: digits = "2"
Output: ["a","b","c"]
*/

package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	//mapping of digit to letterss
	digitMap := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	//start with an empty combination
	result := []string{""}

	//for each digit in the input
	for i := 0; i < len(digits); i++ {
		digit := digits[i]
		letters := digitMap[digit]
		var temp []string

		//combination each exiting resutl with each new letter
		for _, combination := range result {
			for _, letter := range letters {
				temp = append(temp, combination+letter)
			}
		}

		//update the result with the new combination
		result = temp
	}

	return result

}

func main() {
	fmt.Println(letterCombinations("23"))
}
