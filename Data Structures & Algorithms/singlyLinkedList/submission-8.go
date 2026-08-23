type NodeList struct {
    val int
    next *NodeList
}

type LinkedList struct { 
    Head *NodeList
    Tail *NodeList
}

func NewLinkedList() *LinkedList {
    return &LinkedList{}
}

func (ll *LinkedList) Get(index int) int {
    current := ll.Head
    for i := 0; current != nil; i++{
        if i == index {
            return current.val
        }
        current = current.next
    }
    return -1
}

func (ll *LinkedList) InsertHead(val int) {
    if ll.Head == nil && ll.Tail == nil {
        newNode := &NodeList{
            val: val,
            next: nil,
        }
        ll.Head = newNode
        ll.Tail = newNode
        return
    }

    newNode := &NodeList{
        val: val,
        next: ll.Head,
    }

    ll.Head = newNode
}

func (ll *LinkedList) InsertTail(val int) {
    newNode := &NodeList{
        val: val,
        next: nil,
    }
    if ll.Head == nil && ll.Tail == nil {
        ll.Head = newNode
        ll.Tail = newNode
        return
    }

    ll.Tail.next = newNode
    ll.Tail = newNode
}

func (ll *LinkedList) Remove(index int) bool {
    // handle negative case, input less than 0, or linked list is empty
    if index < 0 || ll.Head == nil {
        return false
    }

    // happy path 1, input index 0
    if index == 0 {
        ll.Head = ll.Head.next
        if ll.Head == nil {
            ll.Tail = nil
        }
        return true
    }

    // happy path 2, index > 0, but < tail
    current := ll.Head
    for i := 0; i < index -1 && current != nil; i++{
        if current == nil {
            return false
        }

        // get previous node before index
        current = current.next
    }

    // revalidate negative case, index > linkedlist length
    if current == nil || current.next == nil {
        return false // index out of range
    }

    // removal of ith node
    current.next = current.next.next

    // if index == tail node
    if current.next == nil {
        ll.Tail = current
    }

   
    return true
}

func (ll *LinkedList) GetValues() []int {
    arr := []int{}
    current := ll.Head
    for current != nil {
        arr = append(arr, current.val)
        current = current.next
    }

    return arr
}
