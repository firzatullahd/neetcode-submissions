type MyStack struct {
    q1 []int
    q2 []int
}

func Constructor() MyStack {
    return MyStack{
        q1:[]int{},
        q2:[]int{},
    }
}

func (this *MyStack) Push(x int) {
    this.q1 = append(this.q1, x)
}

func (this *MyStack) Pop() int {
    output := this.Top()
    for i := range this.q1 {
        if i == len(this.q1) - 1 {
            continue
        }

        this.q2 = append(this.q2, this.q1[i])
       
    }
    this.q1 = this.q2[:]
    this.q2 = []int{}
    return output
}

func (this *MyStack) Top() int {
    if this.Empty() {
        return -1
    }
    last := len(this.q1) -1
    return this.q1[last]
}

func (this *MyStack) Empty() bool {
    fmt.Printf("%+v \n", this)
    return len(this.q1) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
