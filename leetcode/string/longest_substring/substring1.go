//Given a string s, find the length of the longest substring
//without repeating characters.

package main 

import(
	"fmt"

)

func lengthOfLongestSubstring(s string)int{

	myMap := make(map[rune]int)
	maxLen := 0
	left := 0

	
	for right , char := range s {

		//Check for Repeating Characters
		if index , found := myMap[char]; found && index >= left{

			// move the left pointer to avoid repetition
			left = index + 1
		}

		// store update the index of the character
		myMap[char] = right

		//update the max length
		if right-left+1 > maxLen{
			maxLen = right - left + 1
		} 
	}

	return maxLen
	
}

func main(){

	s := "pwwkew"
	
	fmt.Printf("The length of the string is %d",lengthOfLongestSubstring(s))
}

/*
Sliding Window Approach for efficiency.
Hash Map (map[rune]int) for fast lookups.
Handles Edge Cases:

"" → Output: 0
"aaaaa" → Output: 1
"abcdef" → Output: 6
*/
