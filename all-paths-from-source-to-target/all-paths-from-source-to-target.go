import (
    "fmt"
)

func allPathsSourceTarget(graph [][]int) [][]int {
    n := len(graph)
    res := [][]int{}
    var dfs func(path []int, src, nei, dst int)
    dfs = func(path []int, src, nei, dst int) {
        nodeSoFar := nei
        if nodeSoFar == dst {
            tmp := make([]int, len(path))
            copy(tmp, path)
            res = append(res, tmp)
            return
        }
        neighbors := graph[nodeSoFar]
        for _, nextNei := range neighbors {
            //if !visited[strPath] {
                path = append(path, nextNei)
                dfs(path, nodeSoFar, nextNei, dst)
                path = path[:len(path)-1]
            }
    }
    src, dst := 0, n-1
    neighbors := graph[0]
    for _, nei := range neighbors {
        dfs([]int{src, nei}, src, nei, dst)
    }
    return res
}
