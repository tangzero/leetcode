package main

// @leet start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := new(ListNode)
	current := sum

	n1 := l1
	n2 := l2

	carry := 0

	for {
		v1 := 0
		if n1 != nil {
			v1 = n1.Val
		}
		v2 := 0
		if n2 != nil {
			v2 = n2.Val
		}

		current.Val = (v1 + v2 + carry) % 10
		carry = (v1 + v2 + carry) / 10

		if n1 != nil {
			n1 = n1.Next
		}
		if n2 != nil {
			n2 = n2.Next
		}

		if n1 == nil && n2 == nil && carry == 0 {
			break
		}

		current.Next = new(ListNode)
		current = current.Next
	}

	return sum
}

// @leet end

