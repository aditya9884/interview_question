//Integer to Roman

package main 

import(
	"fmt"
	"strings"
	"strconv"
)

func intoRoman(num int)string {

	maxRomanNumber := 3999

	// check num is greater, if greater than return integer to string
	if num > maxRomanNumber {
		return strconv.Itoa(num)
	}

	// core lookup table struct
	conversions := []struct{
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions{
		for num >= conversion.value{
			roman.WriteString(conversion.digit)
			num -= conversion.value
		}
	}

	return roman.String()

}

func main(){

	fmt.Println("The resulted roman number is:",intoRoman(66))
}