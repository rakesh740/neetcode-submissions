func climbStairs(n int) int {
    if n == 1 {
        return 1
    } else if n == 2 {
        return 2
    } 
     a, b , c := 1, 2, 0
     for i := 3 ; i <= n ; i++ {
        c = a + b 
        a, b = b, c
     }


    return c
}
