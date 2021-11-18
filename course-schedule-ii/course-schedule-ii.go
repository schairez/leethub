
type Status uint

const (
    Unvisited Status = iota
    Visiting 
    Visited 
)

var (
    ErrCycleDetected = errors.New("")
)


type Graph struct {
    Adj  [][]int
}

func NewGraph(prereqs [][]int, n int) *Graph {
    adj := make([][]int, n)
    for i := range adj {
        adj[i] = make([]int, 0, n)
    }
    for i := range prereqs {
        dst, src := prereqs[i][0], prereqs[i][1]
        adj[src] = append(adj[src], dst)
    }
    return &Graph{adj}
}

func findOrder(numCourses int, prerequisites [][]int) []int {
    n := len(prerequisites)
    if n == 0 {
        res := make([]int, numCourses)
        for i := 0; i < numCourses; i++ {
            res[i] = i
        }
        return res
    }
    visited := make(map[int]Status, numCourses)
    graph := NewGraph(prerequisites, numCourses)
    var err error
    res, err := TopSort(graph, visited, numCourses)
    if err == ErrCycleDetected { 
        return []int{} 
    }
    revSlice(&res)
    
    return res
} 
    
func revSlice( arr *[]int) {
    for i, j := 0, len(*arr)-1; i < j; i, j = i+1, j-1 {
        (*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
    }
}



func TopSort(g *Graph, visited map[int]Status, n int) ([]int, error) {
    res := make([]int, 0, n)
    var dfs func(node int) error
    dfs = func(node int) error {
        visited[node] = Visiting 
        for _, nei := range g.Adj[node] {
            switch visited[nei] {
            case Unvisited:
                if err := dfs(nei); err != nil { 
                    return err 
                }
            case Visiting:
                return ErrCycleDetected
            case Visited:    
                continue
            }
        }
        visited[node] = Visited
        res = append(res, node)
        return nil
    }
    for node := range g.Adj {
        if visited[node] == Unvisited {
            if err := dfs(node); err != nil {
                return nil, err
            }
        }
    }
    return res, nil
}


//prerequisites[i] = [ai, bi]
// bi = src; ai =dst; bi -> ai

/*

Return the ordering of courses you should take to finish all courses. If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.

ex 2
Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
[0] -> [1, 2]
[1] -> [3]
[2] -> [3]

        0
       / \
      1   2
       \ /
        3

Output: [0,2,1,3]

Constraints:

1 <= numCourses <= 2000
0 <= prerequisites.length <= numCourses * (numCourses - 1)
prerequisites[i].length == 2
0 <= ai, bi < numCourses
ai != bi
All the pairs [ai, bi] are distinct.

*/
