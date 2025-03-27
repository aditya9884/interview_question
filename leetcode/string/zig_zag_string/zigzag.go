/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
P   A   H   N
A P L S I I G
Y   I   R
*/

package main 

import(
	"fmt"

)

func converts(s string, numRows int) string{

	//handle the edge case
	if numRows == 1 || numRows >= len(s){
		return s
	}

	//create an arry to hold rows
	rows := make([]string, numRows)

	currentRows := 0
	goingDown := false

	// iterate over each character in the string
	for _, char := range s {

		//append character to the current row
		rows[currentRows] += string(char) 

		//change the direction if at the top or bottom row
		if currentRows == 0{
			goingDown = true
		}else if currentRows == numRows -1 {
			goingDown = false
		}

		//move to the next row based on direction 
		if goingDown{
			currentRows ++
		}else{
			currentRows --
		}
	}

	//concatenate all row to form the final string
	resutl := ""
	for _, row := range rows{
		resutl += row
	}
	return resutl
}


func main(){

	input := "PAYPALISHIRING"
	numRows := 3
	output := converts(input,numRows)
	fmt.Println(output)
}


