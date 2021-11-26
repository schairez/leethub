//topSort Kahn's bfs
//time: O(V+E); space: O(V+E)

func findOrder(numCourses int, prerequisites [][]int) []int {
    graph := make([][]int, numCourses)
    numVisitNodes := 0
    inDegree := make([]int, numCourses)
    queue := []int{}
    for i := range prerequisites {
        dst, src := prerequisites[i][0], prerequisites[i][1]
        graph[src] = append(graph[src], dst)
        inDegree[dst]++
    }
    for nodeID := range inDegree {
        if inDegree[nodeID] == 0 {
            queue = append(queue, nodeID)
        }
    }
    if len(queue) == 0 { return []int{} }
    
    res := make([]int, 0, len(prerequisites))
    for len(queue) != 0 {
        numVisitNodes++
        var nodeID int
        nodeID, queue = queue[0], queue[1:]
        res = append(res, nodeID)
        neighbors := graph[nodeID]
        for _, nei := range neighbors {
            inDegree[nei]--
            if inDegree[nei] == 0 {
                queue = append(queue, nei)
            }
        }
    }
    if numCourses == numVisitNodes {
        return res
    }
    return []int{}
}




//////////////////////////////////////////

//topSort postorder dfs 

type Status uint

const (
    Unvisited Status = iota
    Visiting 
    Visited 
)

var (
    ErrCycleDetected = errors.New("detected cycle")
)


type Graph struct {
    Adj  [][]int
}

func NewGraph(prereqs [][]int, n int) *Graph {
    adj := make([][]int, n)
    for i := range prereqs {
        dst, src := prereqs[i][0], prereqs[i][1]
        adj[src] = append(adj[src], dst)
    }
    return &Graph{adj}
}

func findOrderDFS(numCourses int, prerequisites [][]int) []int {
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
    //rev Slice
    for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
        res[i], res[j] = res[j], res[i]
    }
    
    return res
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
