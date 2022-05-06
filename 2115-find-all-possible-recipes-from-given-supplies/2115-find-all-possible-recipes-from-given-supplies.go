
// Kahn's BFS
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
    n := len(recipes)
    inDegree := make(map[string]int, n)
    recipeSet := make(map[string]struct{}, n)
    for i := range recipes {
        recipeSet[recipes[i]] = struct{}{}
    }
    graph := make(map[string][]string)
    for i := range recipes {
        inDegree[recipes[i]] += len(ingredients[i])
        for _, ingredient := range ingredients[i] {
            if _, exists := graph[ingredient]; !exists {
                graph[ingredient] = make([]string, 0)
            }
            graph[ingredient] = append(graph[ingredient], recipes[i])
        }
    }
    var node string
    queue := make([]string, len(supplies))
    // indegree = 0 for supplies
    for i := range supplies {
        queue[i] = supplies[i]
    }
    for node := range inDegree {
        if inDegree[node] == 0 {
            queue = append(queue, node)
        }
    }
    res := make([]string, 0, n)
    
    // while not empty
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            // dequeue, popleft
            node, queue = queue[0], queue[1:]
            if _, exists := recipeSet[node]; exists {
                res = append(res, node)
            }
            neighbors := graph[node]
            if len(neighbors) != 0 {
                for _, neiNode := range neighbors {
                    inDegree[neiNode]--
                    if inDegree[neiNode] == 0 {
                        queue = append(queue, neiNode)
                    }
                }
            }
        }
    }
    return res
}