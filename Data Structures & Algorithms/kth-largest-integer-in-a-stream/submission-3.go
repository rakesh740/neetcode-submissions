type KthLargest struct {
    cache *IntHeap
	k int
}

type IntHeap []int

func (l IntHeap) Len() int {
	return len(l)
}
func (l IntHeap) Less(i, j int) bool {
	return l[i] < l[j]
}

func (l IntHeap) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

func (l *IntHeap) Push(x any) {
	*l = append(*l, x.(int))
}

func (l *IntHeap) Pop() any {
	old := *l
	n := len(old)
	x := old[n-1]
	*l = old[:n-1]
	return x
}


func Constructor(k int, nums []int) KthLargest {
	h := IntHeap(nums[:])
	heap.Init(&h)
    return KthLargest{
		cache: &h,
		k : k,
	}
}


func (this *KthLargest) Add(val int) int {
	heap.Push(this.cache, val)
	
	for this.cache.Len() > this.k {
	 heap.Pop(this.cache)	
	}
	return (*this.cache)[0]
}









