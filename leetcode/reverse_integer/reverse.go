/*
Given a signed 32-bit integer x, return x with its digits reversed. 
If reversing x causes the value to go outside the signed 32-bit integer
 range [-2^31, 2^31 - 1], then return 0.
Input: x = 123
Output: 321
*/

package main 

import(
	"fmt"
	"math"

)


func reverse(x int)int{

	reversed := 0
	isNegative := x < 0

	if isNegative{
		//convert negative numb  to positive for processing
		x = -x
	}

	for x > 0{
		digit := x % 10
		x /= 10

		// check overflow before adding the digit
		if reversed > (math.MaxInt32 - digit)/10{

			//return 0 if overflow occurs
			return 0
		}

		reversed = reversed * 10 + digit
	}

	if isNegative{
		return - reversed
	}
	return reversed
}

func main(){
	
	fmt.Print(reverse(123))
}

