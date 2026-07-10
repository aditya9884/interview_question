/*
Given an array of integers nums and an integer target, return indices of the
two numbers such that they add up to target.
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
*/
// using the time complexity O(n^2)

package main 

import "fmt"

func twoSum(nums []int, target int)[] int {
	n := len(nums)
	for i:= 0; i < n; i++{
		for j:= i+1; j < n; j++{
			if nums[i] + nums[j] == target{
				return []int{i,j}
			}
		} 
	} 
	return []int{} 
}

func main(){
	b := []int{2,7,11,15}
	c := 9
	fmt.Println(twoSum(b,c))
}
