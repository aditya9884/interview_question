//Given an array of integers nums and an integer target, 
//return indices of the two numbers such that they add up to target.

package main 

import(
	"fmt"
)

func twoSum(nums []int,target int)[]int{
	
	//cerate a map to store numbers and their indices
	myMap:= make(map[int]int)
	for i, num:= range nums{

		complement := target - num // calculate the complement

		//check if the complement exits in the map 
		if idx, found := myMap[complement]; found{

			return []int{idx, i} // return the indicies of the two numbers
		}
		myMap[num] = i // store the index of the current numbers
	}
	
	return nil // return nil if no pair is found 

}



func main(){

	nums := []int{2,7,11,15}
	target := 13
	result := twoSum(nums, target)
	fmt.Println(result)

}
