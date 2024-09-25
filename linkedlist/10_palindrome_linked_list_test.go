package linkedlist

import (
	"fmt"
	"testing"
)

func TestIsPalindromeLinkedList(t *testing.T) {

	var linkedlist SingleLinkedList = SingleLinkedList{
		head: nil,
		tail: nil,
	}

	linkedlist.insert(10)
	linkedlist.insert(20)
	linkedlist.insert(30)

	linkedlist.printLinkedList()

	isPalindrome := linkedlist.isPalindrome()
	fmt.Println("is palindrome:", isPalindrome)

	linkedlist = SingleLinkedList{
		head: nil,
		tail: nil,
	}

	linkedlist.insert(10)
	linkedlist.insert(20)
	linkedlist.insert(30)
	linkedlist.insertAtTail(30)
	linkedlist.insertAtTail(20)
	linkedlist.insertAtTail(10)

	linkedlist.printLinkedList()

	isPalindrome = linkedlist.isPalindrome()
	fmt.Println("is palindrome:", isPalindrome)

}
