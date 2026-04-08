func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)

	for _, v := range nums {
		heap.Push(h, v) 
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return heap.Pop(h).(int)

}
type IntHeap []int

func (a IntHeap) Len() int {
	return len(a)
}
func (a IntHeap) Less(i, j int) bool {
	return a[i] < a[j]
}

 func (a IntHeap) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

 func (a *IntHeap) Push ( val interface{}) {
	*a = append(*a, val.(int))
 }

 func (a *IntHeap) Pop () interface{} {
	old := *a
	n := len(old)
	x := old[n-1]
	*a = old[:n-1]
	return x
 }

