//n*n matrcityAx
//time: O(n*n) ~ O(n^2)
//space: O(n)

func findCircleNum(isConnected [][]int) int {
    cityVisited := make([]int, len(isConnected))
    providenceCnt := 0
    m := len(cityVisited)
    n := len(isConnected[0])
    var dfs func(cityA int)
    dfs = func(cityA int) {
        for cityB := 0; cityB < n; cityB++ {
            if isConnected[cityA][cityB] == 1 && cityVisited[cityB] == 0 {
                cityVisited[cityB] = 1
                dfs(cityB)
                
            }
        }
    }
    for cityA :=0; cityA < m; cityA++ { 
        if cityVisited[cityA] == 0 {
            dfs(cityA)
            providenceCnt++
        }
  
    }    
    
    return providenceCnt
    
}