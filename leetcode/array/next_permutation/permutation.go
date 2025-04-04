/*
A permutation of an array of integers is an arrangement of its members 
into a sequence or linear order.
For example, for arr = [1,2,3], the following are all the permutations of arr: 
[1,2,3], [1,3,2], [2, 1, 3], [2, 3, 1], [3,1,2], [3,2,1]
*/

package main 

import(
	"fmt"
)

func nextPermutation(num []int) []int{

	n := len(num)
	if n <= 1 {
		return num
	}

	i := n-2
	for i >= 0 && num[i] >= num[i+1]{
		i--
	}

	if i >= 0 {

		j := n-1
		for j >= 0 && num[j] <= num[i] {
			j--
		}
		num[i], num[j] = num[j], num[i]
	}

	reverse(num[i+1:])
	return num

}


func reverse(num []int){

	for i , j :=0, len(num) -1; i < j; i, j = i +1, j -1 {
		num[i], num[j] = num[j], num[i]
	}
}


func main(){
	num := []int{1,2,3}

	fmt.Println(nextPermutation(num))
}
