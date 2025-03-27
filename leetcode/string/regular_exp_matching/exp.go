/*
Given an input string s and a pattern p, implement regular expression 
matching with support for '.' and '*' where:

'.' Matches any single character.
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).
*/

package main 

import(

	"fmt"
)


func isMatch(s string, p string) bool{

	memo := make(map[[2]int]bool)
	return dp(0,0,s,p, memo)
}

func dp(i int , j int, s string, p string, memo map[[2]int]bool)bool{

	//memoization check(dp concept)
	if val , ok := memo[[2]int{i,j}]; ok {
		return val
	}

	//base case: End of Pattern(Determines if we've reached the end of the pattern string p.)
	if j == len(p){
		return i == len(s)
	}

	//first charcter match check(Determines if the current characters in s and p match.)
	firstMatch := i < len(s) && (p[j] == s[i] || p[j] == '.')

	//Declares a variable ans to store the result of the current computation.
	var ans bool

	if (j + 1) < len(p) && p[j+1] == '*'{
		ans = dp(i, j+2, s, p, memo) || (firstMatch && dp(i+1, j, s, p, memo))
	}else{
		ans = firstMatch && dp(i+1, j+1, s, p, memo)
	}

	memo[[2]int{i,j}] = ans
	return ans
}

func main(){

	fmt.Print(isMatch("aa","a*"))
}