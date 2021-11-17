//time: O(2^N * N)
//space : O(2^N * N)

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
