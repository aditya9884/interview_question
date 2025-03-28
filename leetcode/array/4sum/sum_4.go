/*
Given an array nums of n integers, return an array of all the unique quadruplets 
[nums[a], nums[b], nums[c], nums[d]] such that:

0 <= a, b, c, d < n
a, b, c, and d are distinct.
nums[a] + nums[b] + nums[c] + nums[d] == target
You may return the answer in any order.
Example 1:
Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
*/


package main 

import(
	"fmt"
	"sort"
)

func fourSum(num []int, target int)[][]int{

	sort.Ints(num)
	result := [][]int{}

	//store the length of input array 
	n := len(num)

	//outer two loop (first two element)
	for i := 0; i < n-3; i++{
		// to skip the duplicate
		if i > 0 && num[i] == num[i-1]{
			continue
		}

		//loop for the second element   
		for j:= i +1; j < n-2; j++{
			//skip the duplicate
			if j > i+1 && num[j] == num[j-1]{
				continue
			}

			//two pointer technique for the last two element
			left, right := j+1 , n-1

			//using while loop 
			for left < right{
				//calculating the sum of four number
				sum := num[i] + num[j] + num[left] + num[right]

				//calculating the sum and add to result
				if sum == target{
					result = append(result, []int{num[i],num[j],num[left],num[right]})

					//skip the duplicat
					for left < right && num[left] ==  num[left+1]{
						left++
					}
					

					// skip the duplicate
					for left > right && num[right] == num[right-1]{
						right --
					}
					left ++
					right --
				}else if sum < target{
					left ++
				}else{
					right --
				}
			}
		}
	}
	return result

	
}

func main(){
	num := []int{-2,-1,-1,1,1,2,2}
	target := 0

	fmt.Println(fourSum(num, target))
}