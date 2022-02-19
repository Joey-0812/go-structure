package array

type Array struct {
	Data []int
	size int
}

// New 构造函数，传入数组的容量capacity构造Array
func New(capacity int) *Array {
	// size 默认为 0
	return &Array{
		Data: make([]int, capacity),
		size: 0,
	}
}

// GetCapacity 获取数组的容量
func (a *Array) GetCapacity() int {
	return len(a.Data)
}

// GetSize 获得数组中的元素个数
func (a *Array) GetSize() int {
	return a.size
}

// IsEmpty 返回数组是否为空
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

func (a *Array) Add(index int, value int) {
	if a.size == len(a.Data) {
		panic("add failed, array is full")
	}

	if index < 0 || index > a.size {
		panic("add failed, index out of range")
	}

	for i := a.size - 1; i >= index; i-- {
		a.Data[i+1] = a.Data[i]
	}

	a.Data[index] = value
	a.size++
}

// AddLast 向所有元素后添加一个新元素
func (a *Array) AddLast(value int) {
	a.Add(a.size, value)
}

// AddFirst 向所有元素前添加一个新元素
func (a *Array) AddFirst(value int) {
	a.Add(0, value)
}
