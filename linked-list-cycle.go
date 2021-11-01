package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	if head.Next == nil {
		return false
	}

	node := head.Next
	mem := map[*ListNode]bool{}
	for {

		if _, ok := mem[node]; ok {
			return true
		} else {
			mem[node] = true
		}
		if node.Next == nil {
			return false
		}

		node = node.Next
	}
}
