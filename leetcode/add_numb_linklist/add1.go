//You are given two non-empty linked lists representing two non-negative integers. 
//The digits are stored in reverse order, and each of their nodes contains a single digit.
//Add the two numbers and return the sum as a linked list.

package main 

import(
	"fmt"
)

type ListNode struct{
	val int 
	next *ListNode
}

type LinkedList struct{
	head *ListNode
	length int 
	
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode{
	// create a dummy node for the result
	dummy := &ListNode{}
	current := dummy

	//initialize cary 
	carry := 0

	//iterate until both list are empty and carry is 0
	for l1 != nil || l2 != nil || carry > 0{
		
		// get the value from eaxh list (or 0 if empty)
		sum := carry
		if l1 != nil{
			sum += l1.val //access val directly
			l1 = l1.next
		}
		if l2 != nil{
			sum += l2.val //access val directly
			l2 = l2.next
		}

		//calculate the digit and carry 
		carry = sum/10
		digit := sum % 10

		//create new node for the digit 
		newNode := &ListNode{val:digit}

		//Append the new node to the result
		current.next = newNode
		current = current.next 
	}
	//return the next of dummy head which is the actual result
	return dummy.next

}

// printlist print the linked list starting from a given node
func PrintList(head *ListNode){
	current := head
	for current != nil{
		fmt.Print(current.val," ")
		current = current.next
	} 
	fmt.Println()
}

func main(){

	//create the first linked list
	l1 := &LinkedList{}
	l1.head = &ListNode{val:7}
	l1.head.next = &ListNode{val:1}
	l1.head.next.next = &ListNode{val:6}

	//create the second list 5,9,2
	l2 := &LinkedList{}
	l2.head = &ListNode{val:5}
	l2.head.next = &ListNode{val:9}
	l2.head.next.next = &ListNode{val:2}

	//add two number
	result := addTwoNumbers(l1.head, l2.head)

	fmt.Println("Result:")

	//print the result which is list node
	PrintList(result)
	
}
