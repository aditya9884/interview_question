/*
Given a roman numeral, convert it to an integer.
Input: s = "III"
Output: 3
Explanation: III = 3.
*/

package main 

import(

	"fmt"
)

func romanToInts(s string) int{

	romanMap := map[byte]int{

		'M' : 1000,
		'D' : 500,
		'C' : 100,
		'L' : 50,
		'X' : 10,
		'V' : 5,
		'I' : 1,
	}
	

	total := 0
	preValue := 0

	for i:= len(s) -1; i >= 0; i--{
		current := romanMap[s[i]]

		if current < preValue {
			total -= current
		}else{
			total += current
		}
		preValue = current
	}
	return total

}


func main(){
	fmt.Println(romanToInts("LVIII"))
}

