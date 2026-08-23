/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var prev, current, next *ListNode

	current = head
	for current != nil {
		// save next node to a var, before overwrite
		next = current.Next

		// repoint current node to previous node
		current.Next = prev

		// save current node as previous (if last iteration, last node becomes prev)
		prev = current

		// focus to next node
		current = next
	}

	return prev
}
