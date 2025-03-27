//Given a string s, return the longest palindromic substring in s.

 
package main 

import(
	"fmt"

)


func longestPalindrome(s string)string {
	if len(s) < 1 {
		return ""
	}

	start, end := 0 , 0

	for i :=0; i< len(s); i++{
		// odd length of the palidrome
		len1 := expandAroundcenter(s,i,i)

		// even length of the palidrome
		len2 := expandAroundcenter(s,i,i+1) 
		maxLen := max(len1,len2)

		if maxLen > end - start {
			start = i -(maxLen -1)/2
			end = i + maxLen/2
		}
	}
	return s[start : end + 1]
}

func expandAroundcenter(s string, left int, right int)int{
	for left >= 0 && right < len(s) && s[left] == s[right]{
		left --
		right ++
	}
	return right - left -1
}

func max(a, b int)int{
	if a > b{
		return a
	}
	return b
}


func main(){

	s := "babad"
	fmt.Print(longestPalindrome(s))
	
}

