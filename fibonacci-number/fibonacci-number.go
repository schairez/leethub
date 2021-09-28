/*                          
                    fib(4) 3 
                 2  /     \   1
               fib(3)  +  fib(2)
           1 /       \  1 /     \ 0
        fib(2)  + fib(1) fib(1)  + fib(0)
        / 1    0\                      
      fib(1) + fib(0)
*/

func fib(n int) int {
    memo := map[int]int{}
    return helper(n, memo )
}


func helper(n int, memo map[int]int) int {
    if n == 1 || n == 0 {
        return n
    }
    if _, ok := memo[n]; !ok {
        memo[n] = helper(n-1, memo) + helper(n-2, memo)
    }
    return memo[n]
    //var lSum, rSum int 
    //                f(3)
    //          f(2, 0) + f(1, 0)
    //       f(1) + f(0)    
    //helper(n-1, currSum + n ) + helper(n-2, currSum + )
    
}

