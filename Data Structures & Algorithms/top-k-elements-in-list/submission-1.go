func topKFrequent(nums []int, k int) []int {
	h := &elems{}
	heap.Init(h)
	cache := make(map[int]int)
	for _, v := range nums {
		cache[v]++
	}

	for k, v := range cache {
		heap.Push(h, elem{
			num: k,
			freq: v,
		})
	}
	
	var res []int

	for i:=k; i > 0 ;i-- {
		e := heap.Pop(h).(elem)
		res = append(res, e.num)
	}

return res

}

type elem struct {
	num int
	freq int
}

type elems []elem

func (e elems) Len () int {return len(e)}
func (e elems) Less (i, j int) bool {return e[i].freq > e[j].freq }
func (e elems) Swap (i, j int)  {e[i], e[j] = e[j] , e[i]}

func(e *elems) Pop () (x interface{}) {
	x = (*e)[len(*e)-1]
	*e = (*e)[0:len(*e)-1]
	return
}

func(e *elems) Push (x interface{}) {
	*e = append(*e, x.(elem))
}