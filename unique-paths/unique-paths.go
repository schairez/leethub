/*
time: O(m*n) number of nodes in graph
space: O(n) for aux memo
*/


import "strconv"


func uniquePaths(m int, n int) int {
    memo := make(map[string]int)
    return dfs(memo, 0, 0, m, n )
}


func isValid(i, j, m, n int) bool {
    if i == m || j == n {
        return false
    }
    return true 
}

func dfs(memo map[string]int, i, j, m, n int) int {
    if !isValid(i,j,m,n) {
        return 0
    }
    key := strconv.Itoa(i) + "_" + strconv.Itoa(j)
    if _, ok := memo[key]; ok {
        return memo[key]
    }
    if i == m -1 && j == n - 1 {
        return 1 
    }
    path := 0 
    path += dfs(memo,i+1, j, m, n)
    path  += dfs(memo,i, j+1, m, n)
    memo[key] = path
    
    return memo[key]
    
}

//dfs helper; preorder(DLR)


/*
down or right           [0,0]
                    dfs(0,0)
                   /
               dfs(1,0)
               
*/





























/*
m=3; n = 2
down, right
                    start ->[0][0]
                    end -> [m-1][n-1]
                   
                   down -> [currRow+1][0]
                   //down terminating -> when currRow == m-1
                   |-----v
                   | ----v
                   |---- e_
     3x2
       visited[m][n] = [[0, 0 ]
                        [0, 0]
                  "R0C0"
        dfs([0,0],str, memo)
            if memo[str] exists
               return
            if i==m-1 && j == n-1: //
                res++ 
            if isValid(i,j) //f(2,0)
                dfs([i+1,j]) //down
                dfs([i,j+1]) //right
            
             (0,0)
              /
             (1,0)          memo[1][0] = 1 
             /    \
            (2,0)  (1,1)
            /    \ 
          (3,0)  (2,1)
          
             memo[i][j] ---> memo[i][j]       R0C0
            path = "R0C0R1C0R2C0R2C1" 
            map[str]bool
            memo
          
          
            
  func isValid()
     if row != m && if col != n
        
*/

