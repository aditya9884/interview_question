/*
Given an integer array nums and an integer val, remove all occurrences of val in nums
in-place. The order of the elements may be changed. Then return the number of
elements in nums which are not equal to val.
Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 2.
It does not matter what you leave beyond the returned k (hence they are underscores).
*/

package main

import (
	"fmt"
)

func removeElement(num []int, val int) int {

	if len(num) == 0 {
		return 0
	}

	k := 0

	for _, nums := range num {

		if nums != val {

			num[k] = nums
			k++
		}
	}
	return k
}

func main() {
	num := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2

	fmt.Println(removeElement(num, val))
	fmt.Println("New length", removeElement(num, val))
	fmt.Println("Modified array :", num[:removeElement(num, val)])
}
