//Given two sorted arrays nums1 and nums2 of size m and n respectively, 
//return the median of the two sorted arrays.

package main

import(
	"fmt"
	"math"
)

func findMedianSortedArray(num1 []int, num2 []int) float64{

	// ensure num1 is smaller than num2
	if len(num1) > len(num2){
		num1, num2 = num2, num1
	}

	m, n := len(num1), len(num2)
	halfLen := (m + n + 1)/2
	low , high := 0, m

	for low <= high{
		partX := (low + high) / 2
		partY := halfLen - partX

		maxX := math.MinInt32
		if partX > 0{
			maxX = num1[partX-1]
		}

		minX := math.MaxInt32
		if partX < m{
			minX = num1[partX]
		} 

		maxY := math.MinInt32
		if partY > 0{
			maxY = num2 [partY - 1]
		}

		minY := math.MaxInt32
		if partY < n{
			minY = num2[partY]
		}

		if maxX <= minY && maxY <= minX {
			if (m+n)%2 == 0 {
				return float64(max(maxX,maxY) + min(minX,minY))/2.0
			}else{
				return float64(max(maxX,maxY))
			}
		}else if maxX > minY {
			high = partX - 1
		}else{
			low = partX + 1
		}
	}
	return 0.0
}

func max(a,b int)int{
	if a> b{
		return a
	}
	return b
}

func min(a,b int)int{
	if a < b{
		return a
	}
	return b
}

func main(){

	num1 := []int{1,2}
	num2 := []int{3,4}

	fmt.Println(findMedianSortedArray(num1,num2))
}

