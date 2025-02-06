/*
Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer.
Input: s = "42"
Output: 42
*/

package main 

import(
	"fmt"
	"math"
	"unicode"
)

func myAtoi(s string)int {

	n := len(s)
	if n == 0 {
		return 0
	}

	// trim the leading white space
	i := 0
	for i < n && s[i] == ' '{
		i++
	}

	//check if empty after trimming 
	if i == n {
		return 0
	}

	//check sign
	sing := 1
	if s[i] == '+'{
		i ++
	}else if s[i] == '-'{
		sing = -1
		i++
	}

	//convert digit
	result := 0
	for i < n && unicode.IsDigit(rune(s[i])){
		digit := int(s[i]-'0')
		//check overflow
		if result > (math.MaxInt32 - digit)/10{
			if sing == 1 {
				return math.MaxInt32
			}else{
				return math.MinInt32
			}
		}
		result = result * 10 + digit
		i++
	}
	return result * sing 
}


func main(){

	fmt.Print(myAtoi("0-9"))
}

