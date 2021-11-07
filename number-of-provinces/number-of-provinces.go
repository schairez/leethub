//time: O(N^2)
//space: O(N)

func findCircleNum(isConnected [][]int) int {
    isVisited := make([]int, len(isConnected))
    provinceCnt := 0 
    m := len(isConnected)
    n := len(isConnected[0])
    var dfs func(i int)
    dfs = func(i int) {
        for j := 0; j < n; j++ {
            if isConnected[i][j] == 1 && isVisited[j] == 0 {
                isVisited[j] = 1
                dfs(j)
            }
        }
    }
    
    for i :=0; i < m; i++ { 
        if isVisited[i] == 0 {
            dfs(i)
            provinceCnt++
        }
  
    }    
    
    return provinceCnt 
    
}