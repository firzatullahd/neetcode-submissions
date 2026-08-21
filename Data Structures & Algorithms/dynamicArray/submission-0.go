type DynamicArray struct {
	items []int
	length int
	capacity int
}

func NewDynamicArray(capacity int) *DynamicArray {
	return &DynamicArray{
		items: make([]int,capacity),
		length: 0,
		capacity: capacity,
	}
}

func (da *DynamicArray) Get(i int) int {
	return da.items[i]
}

func (da *DynamicArray) Set(i int, n int) {
		da.items[i] = n
}

func (da *DynamicArray) Pushback(n int) {
	if da.length == da.capacity {
		da.resize()
	}

	da.items[da.length] = n
	da.length++
}

func (da *DynamicArray) Popback() int {
	if da.length >0 {
		da.length--
	}
	return da.items[da.length]
}

func (da *DynamicArray) resize() {
	da.capacity = da.capacity * 2
	newArr := make([]int, da.capacity)

	for i := 0; i< da.length; i++ {
		newArr[i] = da.items[i]
	}

	da.items = newArr
	
}

func (da *DynamicArray) GetSize() int {
	return da.length
}

func (da *DynamicArray) GetCapacity() int {
	return da.capacity
}
