//time: O(|V| + |E|); in this case O(N) since rooted tree we just visit each node once
//space: O(N)

/*
- each employee has one direct mgr
1 <= n <= 105
0 <= headID < n
*/

type pair struct {
    nodeID, timeSoFar int
}

func max(a, b int) int { if a >= b { return a}; return b}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
    graph := make([][]int, n)
    //idx is emp; value is mgr of emp 
    for emp, mgr := range manager {
        if mgr != -1 {
            graph[mgr] = append(graph[mgr], emp)
        }
    }
    //time needed to inform all employees
    timeNeeded := 0
    //bfs approach
    queue := make([]pair, n)
    queue = append(queue, pair{headID, informTime[headID]})
    var pollNode pair
    for len(queue) != 0 {
        pollNode, queue = queue[0], queue[1:]
        timeNeeded = max(timeNeeded, pollNode.timeSoFar)
        for _, subordinate := range graph[pollNode.nodeID] {
            queue = append(queue, pair{subordinate,
                                       pollNode.timeSoFar + informTime[subordinate]})
        }
    }
    return timeNeeded
}
