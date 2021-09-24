//time: O(n)
//space: O(n)

import "fmt"

type status uint8

//other variations use white,gray,black as indicators of state
const (
    unprocessed status = iota 
    processing
    processed
)

func dfs(adjList [][]int, statusArr []status, node int) error {
    statusArr[node] = processing
    for _, child := range adjList[node] {
        switch statusArr[child] {
            case processing:
                return fmt.Errorf("cycle detected")
                break
            case unprocessed:
                err := dfs(adjList, statusArr, child)
                if err != nil { //dag is false
                    return fmt.Errorf("cycle detected")
                }
            case processed:
                continue
        }
    } 
    statusArr[node] = processed
    return nil 
}        

func canFinish(numCourses int, prerequisites [][]int) bool {
    res := true    //assume dag is true initially
    // bi -> ai; bi is prereq to ai
    graph := make([][]int, numCourses) 
    for i:=0; i < len(prerequisites); i++ {
        prereq, nextCourse := prerequisites[i][1], prerequisites[i][0]
        graph[prereq] = append(graph[prereq], nextCourse)
    }
    // maintain a status slice as we process dfs calls
    statusArr := make([]status, numCourses)
    for i := range statusArr {
        if statusArr[i] != processed {
            err := dfs(graph, statusArr, i)
            if err != nil { //dag is false
                return false
            }
        }
    }
    return res
    

}


/*
//works too; this was a previous version I made 

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	//marked arr
	//
	color := make([]int, numCourses)

	for i := 0; i < len(prerequisites); i++ {
		course, req := prerequisites[i][0], prerequisites[i][1]
		graph[req] = append(graph[req], course)
		//inDegree[course]++
	}
	isDag := true
	var dfs func(node int)
	dfs = func(node int) {
		if !isDag {
			return
		}
		color[node] = 1 //gray
		for _, neighbor := range graph[node] {
			switch color[neighbor] {
			//if node has neighbor which is gray then there's a cycle
			case 1:
				isDag = false
				break
				//white - unvisited so call dfs
			case 0:
				dfs(neighbor)
				//black is marked so skip
			case -1:
				continue
			}
		}
		color[node] = -1

	}
	//do dfs on undiscovered nodes
	for i := range color {
		dfs(i)
	}
	return isDag


}
*/