/*
Find two lines that together with the x-axis form a container, 
such that the container contains the most water.
Return the maximum amount of water a container can store.
You are given an integer array height of length n. There are 
n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
*/

package main 

import(
	"fmt"
)

func maxArea(height []int)int{

	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		//calculate the current area 
		h := min(height[left], height[right])
		width := right - left
		currentArea := h * width

		//update max area if current area is large
		if currentArea > maxArea {
			maxArea = currentArea
		}

		// move the pointer at the shorter line 
		if height[left] < height[right] {
			left++
		}else {
			right--
		}
	}

	return maxArea
}

func main(){

	res := []int{1,8,6,2,5,4,8,3,7}

	fmt.Println("The maximum area is:",maxArea(res))
}

