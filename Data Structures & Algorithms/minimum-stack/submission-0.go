type MinStack struct {
	items []int
}

func Constructor() MinStack {
	return MinStack{items:[]int{}}
}

func (this *MinStack) Push(val int) {
	this.items = append(this.items, val)
}

func (this *MinStack) Pop() {
	this.items = this.items[:len(this.items)-1]
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
	min := this.items[0]
	for i := range this.items {
		if this.items[i] < min {
			min = this.items[i]
		}
	}
	return min
}
