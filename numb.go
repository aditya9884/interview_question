/*
Given an integer x, return true if x is a palindrome, and false otherwise.
Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
*/

package main 

import "fmt"

func isPalindrome(x int) bool{

	if x < 0 {
		return false
	}
	
	// initialise the variable to store the reversed number
	reverse := 0 

	// store the original number
	original := x 

	for x > 0 {
		
		digit := x % 10 // extract the last digit
	
		reverse = (reverse*10) + digit // build the reverse number

		x /= 10 // remove the last number
		
	}

	return original == reverse

}

func main(){
	a := 0
	fmt.Println(isPalindrome(a))
}