//graph coloring dfs v

type color int

const (
    unknown color = 0
    blue color = 1
    red color = -1
)

func possibleBipartition(n int, dislikes [][]int) bool {
    graph := make([][]int, n)
    for i := range graph {
        graph[i] = []int{}
    }
    for _, dislikePair := range dislikes {
        fmt.Println(dislikePair)
        src, dst := dislikePair[0]-1, dislikePair[1]-1
        graph[src] = append(graph[src], dst) 
        graph[dst] = append(graph[dst], src) 
    }
    peopleColor := make([]color, n)
    
    var dfsColoring func(personIdx int, assignColor color ) bool
    dfsColoring = func(personIdx int, assignColor color ) bool {
        peopleColor[personIdx] = assignColor
        for _, nei := range graph[personIdx] {
            if peopleColor[nei] == assignColor {
                return false
            }
            if peopleColor[nei] == unknown && !dfsColoring(nei, -assignColor) {
                return false
            }
        }
        return true
    }
    
    for i:=0; i < n; i++ {
        if peopleColor[i] == unknown && !dfsColoring(i, red) {
            return false
        }
    }
    return true
}





