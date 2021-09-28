
/*                          
fib(6)-> 5+3--> 8
fib(5)-> 5
                    fib(4) 3 
                 2  /     \   1
               fib(3)  +  fib(2)
           1 /       \  1 /     \ 0
        fib(2)  + fib(1) fib(1)  + fib(0)
        / 1    0\                      
      fib(1) + fib(0)
*/
// f(5) i:=n; for i > 2; 
//
//  in: f(5); 
///  i = 5+1 ; n1, n0 = 1,0; 
//for i > 2 
//  i @ 6 { nNew = n1+n0 //1; n0 = n1 //1;  n1 = nNew //1} i--
//  i @ 5 { nNew -> 2; n0 -> 1; n1 = 2} i-- 
//  i @ 4 { nNew -> 3; n0 -> 2; n1 -> 3 }
//  i @ 3 { nNew -> 5; n0 -> 3; n1 -> 5 }
// 


//dp 3 vars optimized
// time: O(n); space:O(1)
func fib(n int) int {
    if n == 1 || n == 0 {
        return n
    }
    var nNew int
    nCurr, nPrev := 1, 0 
    if n == 2 { return nCurr + nPrev } 
    i := n + 1
    for i > 2 {
        nNew = nCurr + nPrev
        nCurr, nPrev = nNew, nCurr 
        i--
    }
    return nCurr
}


//memo top-down approach
// time: O(N); space: O(n) implicit recursive stack and memo table
/*


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

*/
