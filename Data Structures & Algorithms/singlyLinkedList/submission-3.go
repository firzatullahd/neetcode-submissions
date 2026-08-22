type LinkedList struct {
    Length int
    Head *ListNode
}

type ListNode struct {
    Val int
    Next *ListNode
}

func NewLinkedList() *LinkedList {
    return &LinkedList{
        Length: 0,
        Head: nil,
    }
}

func (ll *LinkedList) Get(index int) int {
    if index < 0 || index >= ll.Length {
        return -1
    }
    current := ll.Head
    for i := 0; i < index; i++ {
        current = current.Next
    }
    return current.Val
}

func (ll *LinkedList) InsertHead(val int) {
    newNode := &ListNode{
        Val: val,
        Next: ll.Head,
    }
    ll.Head = newNode
    ll.Length++
}

func (ll *LinkedList) InsertTail(val int) {
    if ll.Head == nil {
        ll.InsertHead(val)
        return
    }
    current := ll.Head
    for current.Next != nil {
        current = current.Next
    }
    current.Next = &ListNode{
        Val: val,
        Next: nil,
    }
    ll.Length++
}

func (ll *LinkedList) Remove(index int) bool {
    if index < 0 || index >= ll.Length {
        return false
    }
    if index == 0 {
        ll.Head = ll.Head.Next
    } else {
        current := ll.Head
        for i := 0; i < index-1; i++ {
            current = current.Next
        }
        current.Next = current.Next.Next
    }
    ll.Length--
    return true
}

func (ll *LinkedList) GetValues() []int {
    arr := []int{}
    current := ll.Head
    for current != nil {
        arr = append(arr, current.Val)
        current = current.Next
    }
    return arr
}