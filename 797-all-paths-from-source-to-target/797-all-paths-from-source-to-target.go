

// BFS approach
// time: O(2^n); and O(n) for copy operation (likely?) ≈ O(n * 2^n)
// space: O(2^n); longest path can store 2^n nodes and O(n) for copy operation ≈ O(n* 2^n) 
// src = 0; sink = n-1
func allPathsSourceTarget(graph [][]int) [][]int {
    n := len(graph)
    src, sink := 0, n-1
    var res [][]int
    queue := make([][]int, 0)
    queue = append(queue, []int{src})
    var path []int
    for len(queue) != 0 {
        path, queue = queue[0], queue[1:]
        dst := path[len(path)-1]
        if dst == sink {
            res = append(res, path)
        }
        for _, nei := range graph[dst] {
            tmp := make([]int, len(path)+1)
            copy(tmp, path)
            tmp[len(tmp)-1] = nei
            queue = append(queue, tmp)
        }
    }
    return res
}
    




















    
    
    
    
    
    
    
    
    

//time: O(2^N * N)
//space : O(2^N * N)
/*
func allPathsSourceTarget(graph [][]int) [][]int {
    n := len(graph)
    res := [][]int{}
    var dfs func(path []int, nei, dst int)
    dfs = func(path []int, nei, dst int) {
        nodeSoFar := nei
        if nodeSoFar == dst {
            tmp := make([]int, len(path))
            copy(tmp, path)
            res = append(res, tmp)
            return
        }
        neighbors := graph[nodeSoFar]
        for _, nextNei := range neighbors {
            path = append(path, nextNei)
            dfs(path, nextNei, dst)
            path = path[:len(path)-1]
        }
    }
    src, dst := 0, n-1
    neighbors := graph[0]
    for _, nei := range neighbors {
        dfs([]int{src, nei}, nei, dst)
    }
    return res
}

*/