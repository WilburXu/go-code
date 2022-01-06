package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	listNode := new(ListNode)
	listNode.Val = 1

	node2 := new(ListNode)
	node2.Val = 2
	listNode.Next = node2

	node3 := new(ListNode)
	node3.Val = 3
	node2.Next = node3

	node4 := new(ListNode)
	node4.Val = 4
	node3.Next = node4

	node5 := new(ListNode)
	node5.Val = 5
	node4.Next = node5

	//reverseList(listNode)
	recursionList := recursion(listNode)
	log.Println(recursionList)
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func recursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := recursion(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}