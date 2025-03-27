/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] 
such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation: 
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
*/

package main 

import(
	"fmt"
	"sort"
	
)

func threeSum(nums []int) [][]int {

	var res [][]int

	sort.Ints(nums)
	

	for i := 0; i < len(nums)-2; i++{

		if nums[i] > 0{
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l , r := i+1, len(nums)-1; l< r; {

			if l > i+1 && nums[l] == nums[l-1]{
				l++
				continue
			}
			if r < len(nums)-1 && nums[r] == nums[r+1]{
				r--
				continue
			}
			switch sum := nums[i] + nums[r] + nums[l]; {

			case sum < 0:
				/* code */
				l++
			case sum > 0:
				r--
			default:
				/* code */
				res = append(res,[]int{nums[i],nums[l],nums[r]})
				l++
				r--
			
			}
		}
	}
	return res
	
}

func main(){

	nums := []int{-1,0,1,2,-1,-4}

	fmt.Println(threeSum(nums))

}

