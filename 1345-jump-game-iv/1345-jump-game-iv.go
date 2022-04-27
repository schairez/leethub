func minJumps(arr []int) int {
    eqi := make(map[int][]int)
    vis := make(map[int]bool)
    n := len(arr)
    dp := make([]int, n)
    for i, v := range arr{
        eqi[v] = append(eqi[v], i)
        dp[i] = 1e9
    }
    queue := make([]int, 1, n)
    queue[0] = 0
    dp[0] = 0
    for len(queue) != 0{
        i := queue[0]
        queue = queue[1:]
        if !vis[arr[i]]{
            for _, j := range eqi[arr[i]]{
                if dp[j] == 1e9{
                    dp[j] = dp[i] + 1
                    queue = append(queue, j)
                }
            }
            vis[arr[i]] = true
        }
        if i + 1 < n && dp[i + 1] == 1e9{
            dp[i + 1] = dp[i] + 1
            queue = append(queue, i + 1)
        }
        if i - 1 > 0 && dp[i - 1] == 1e9{
            dp[i - 1] = dp[i] + 1
            queue = append(queue, i - 1)
        }
    }
    return dp[n - 1]
}


// time: O(n)
// space: O(n)

func minJumps2(arr []int) int {
    n := len(arr)
    if n == 1 {
        return 0
    }
    dst := n-1
    graph := make(map[int][]int, 0)
    for i, v := range arr {
        if _, ok := graph[v]; !ok {
            graph[v] = make([]int, 0)
        }
        graph[v] = append(graph[v], i)
    }
    queue := make([]int, 0, n)
    queue = append(queue, 0)
    visited := make([]bool, n) 
    visited[0] = true
    lvl := 0
    var node int
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            if node == dst {
                return lvl 
            }
            if nei := node + 1; nei < n && !visited[nei] {
                queue = append(queue, nei)
                visited[nei] = true
            }
            if nei := node - 1; nei >= 0 && !visited[nei] {
                queue = append(queue, nei)
                visited[nei] = true
            }
            neighbors, exists := graph[arr[node]] 
            if exists {
                for _, nei := range neighbors {
                    if !visited[nei] {
                        queue = append(queue, nei)
                        visited[nei] = true
                    }
                }
                delete(graph, arr[node])
            }
        }
        lvl++
    }
    return -1
}

/*
func minJumps(arr []int) int {
    n := len(arr)
    if n == 1 {
        return 0
    }
    dst := n-1
    graph := make(map[int][]int, n)
    for i, v := range arr {
        if _, ok := graph[v]; !ok {
            graph[v] = make([]int, 0)
        }
        graph[v] = append(graph[v], i)
    }
    queue := make([]int, 0, n)
    queue = append(queue, 0)
    sameNumberVisited := make(map[int]struct{})
    visited := make([]bool, n) // for visiting nodes
    visited[0] = true
    lvl := 0
    var node int
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            if node == dst {
                return lvl 
            }
            if nei := node + 1; nei < n && !visited[nei] {
                queue = append(queue, nei)
                visited[nei] = true
            }
            if nei := node - 1; nei >= 0 && !visited[nei] {
                queue = append(queue, nei)
                visited[nei] = true
            }
            if _, seenNumber := sameNumberVisited[arr[node]]; !seenNumber {
                sameNumberVisited[arr[node]] = struct{}{}
                neighbors := graph[arr[node]] 
                if len(neighbors) > 1 {
                    for _, nei := range neighbors {
                        if !visited[nei] {
                            queue = append(queue, nei)
                            visited[nei] = true
                        }
                    }
                }
            }
        }
        lvl++
    }
    return -1
}
*/