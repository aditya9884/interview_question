/*
Given an integer array nums sorted in non-decreasing order, remove
the duplicates in-place such that each unique element appears only once.
The relative order of the elements should be kept the same.
Then return the number of unique elements in nums.
Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two elements of nums
being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
*/

package main

import (
	"fmt"
	"strings"
	"strconv"
)

func removeDuplicates(num []int) int {  
    if len(num) == 0 {
		return 0
	}

	// represent the index where the next unique element present 
	k := 0
	for i := 1; i < len(num); i++ {

		if num[i] != num[k] {
			k++
			num[k] = num[i]
		}
		
	}
	return k + 1
	
}

func formateOutput(num []int, k int) string{
	var output []string

	for i := 0; i < len(num); i++{
		if i < k{
			output = append(output, strconv.Itoa(num[i]))
		}else {
			output = append(output, "_")
		}
	}
	return "[" + strings.Join(output, ",") + "]"
}

func main() {

	testcase := [][]int{{1, 1, 2}}

	
	for _, nums := range testcase{
		arr := make([]int,len(nums))
		copy(arr, nums)

		k := removeDuplicates(arr)

		fmt.Printf("k = %d, array = %s\n", k , formateOutput(arr,k))
	}

}
