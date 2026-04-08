func maxProfit(prices []int) int {
    var maxP, lowP int  = math.MinInt32, math.MaxInt32

    for _, v := range prices {
        lowP = min(lowP, v)
        maxP = max(maxP, v - lowP)
    }

    return maxP
}
