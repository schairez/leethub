
func numBusesToDestination(routes [][]int, source int, target int) int {
    graph := buildGraph(routes)
    numRoutes := 0
    visitedStops := make(map[int]struct{})
    visitedRoutes := make(map[int]struct{})
    queue := []int{source}
    var busStop int
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            busStop, queue = queue[0], queue[1:]
            if busStop == target {
                return numRoutes
            }
            //routeIDs := graph[busStop]
            for routeID := range graph[busStop] {
                if _, visitedRoute := visitedRoutes[routeID]; !visitedRoute {
                    for _, busStop := range routes[routeID] {
                        if _, exists := visitedStops[busStop]; !exists {
                            queue = append(queue, busStop)
                            visitedStops[busStop] = struct{}{}
                        }
                    }
                }
                visitedRoutes[routeID] = struct{}{}
            }  
        }
        numRoutes++
    }
    
    return -1
}


func buildGraph(routes [][]int) map[int]map[int]struct{} {
    graph := make(map[int]map[int]struct{})
    for routeID, busStops := range routes {
        for _, busStop := range busStops {
            _, exists := graph[busStop]
            if !exists {
                graph[busStop] = make(map[int]struct{})
            }
            graph[busStop][routeID] = struct{}{}
        }
    }
    return graph
}