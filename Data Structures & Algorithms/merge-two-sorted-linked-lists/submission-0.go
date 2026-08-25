/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	arr := []int{}

	current := list1
	for current != nil {
		arr = append(arr,current.Val)
		current = current.Next
	}

	current = list2
	for current != nil {
		arr = append(arr,current.Val)
		current = current.Next
	} 

	sort.Ints(arr)
	fmt.Println(arr)

	if len(arr) == 0 {
		return nil
	}

	newNode := &ListNode{
		Val: arr[0],
	}
	head := newNode
	for i := 1; i< len(arr); i++ {
		next := &ListNode{
			Val: arr[i],
		}
		newNode.Next = next
		newNode = newNode.Next
	}
	
	return head
}
