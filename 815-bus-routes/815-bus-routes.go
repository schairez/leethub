// stop to route mapping

func numBusesToDestination(routes [][]int, source int, target int) int {
    graph := buildGraph(routes)
    queue := []int{source}
    visitedStops := make(map[int]struct{})
    visitedRoutes := make(map[int]struct{})
    var numBuses int
    var busStop int
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            busStop, queue = queue[0], queue[1:]
            if busStop == target {
                return numBuses
            }
            for routeId := range graph[busStop] {
                if _, visitedRoute := visitedRoutes[routeId]; visitedRoute {
                    continue
                }
                for _, nextBusStop := range routes[routeId] {
                    if _, visitedBusStop := visitedStops[nextBusStop]; visitedBusStop {
                        continue
                    }
                    queue = append(queue, nextBusStop)
                    visitedStops[nextBusStop] = struct{}{}
                }
                visitedRoutes[routeId] = struct{}{}
            } 
        }
        numBuses++
    }
    return -1
}

func buildGraph(routes [][]int) map[int]map[int]struct{} {
    graph := make(map[int]map[int]struct{})
    for routeId, busStops := range routes {
        for _, stop := range busStops {
            if _, exists := graph[stop]; !exists {
                graph[stop] = make(map[int]struct{})
            }
            graph[stop][routeId] = struct{}{} 
        }
    }
    return graph
}