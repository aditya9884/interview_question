/*
Given an integer array nums of length n and an integer target, 
find three integers in nums such that the sum is closest to target.
Return the sum of the three integers. Example:
Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
*/

package main 

import(
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(num []int, target int)int {

	sort.Ints(num)
	//get the sum of first 3 subarray and store the value in closest
	closest := num[0] + num[1] + num[2]

	//outer loop where iterate over the i = 0 
	for i := 0; i < len(num) -2; i++ {

		//two pointer initialization
		left , right := i+1 , len(num) -1

		// inner while loop use to evaluate the diff set of tripplets(subarray)
		for left < right {

			//sum of the current tripplets is cal using element of index i,left, right
			sum := num[i] + num[left] + num[right]

			// compare the closest sum
			if math.Abs(float64(sum - target)) < math.Abs(float64(closest - target)){
				closest = sum
			}

			//adjust pointer 
			if sum < target{
				left ++
			}else if sum > target{
				right --
			}else {
				return sum
			}
		}

	}
	return closest

}

func main(){
	num := []int{-1,2,1,-4}
	target := 2

	fmt.Println(threeSumClosest(num, target))
}

