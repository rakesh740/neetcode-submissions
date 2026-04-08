func minCostClimbingStairs(cost []int) int {
    var mincost = make([]int, len(cost)+1)
	mincost[0] = 0
	mincost[1] =  0
	
	for i:=2; i<len(cost)+1; i++ {
		mincost[i] = min(mincost[i-2] + cost[i-2],  mincost[i-1] + cost[i-1])
	} 

	return mincost[len(cost)]
}
