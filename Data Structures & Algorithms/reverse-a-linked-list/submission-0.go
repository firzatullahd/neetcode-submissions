/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    arr := []*ListNode{}

	current := head
	for current != nil {
		arr = append(arr, current)
		current = current.Next
	}

	if len(arr) == 0 {
		return head
	}

	newHead := arr[len(arr)-1]
	for i := len(arr) -1; i >= 0; i--{
		current = arr[i]
		if i - 1 < 0 {
			current.Next = nil
		} else {
			current.Next = arr[i-1]
		}
		
	}

	return newHead
}
