func lastStoneWeight(stones []int) int {
h := IntHeap(stones)
	he := &h
	heap.Init(he)
	for he.Len() > 1 {
		x1, x2 := heap.Pop(he), heap.Pop(he)
		diff := x1.(int) - x2.(int)
		if diff > 0 {
			heap.Push(he, diff)
		}
	}
	if he.Len() == 0 {
		return 0
	}

	return (*he)[0]
}


type IntHeap []int

func (a IntHeap) Len() int           { return len(a) }
func (a IntHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a IntHeap) Less(i, j int) bool { return a[i] > a[j] }
func (a *IntHeap) Pop() interface{} {
	old := *a
	n := len(old)
	x := old[n-1]
	*a = IntHeap(old[:n-1])
	return x
}

func (a *IntHeap) Push(val interface{}) {
	*a = append(*a, val.(int))
}