

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
    recipeSet := make(map[string]struct{}, len(recipes))
    for i := range recipes {
        recipeSet[recipes[i]] = struct{}{}
    }
    graph := make(map[string][]string)
    // in degree counter for each node
    inDegree := make(map[string]int)
    // O(V+E)
    for i := range recipes {
        inDegree[recipes[i]] += len(ingredients[i])
        for _, ingredient := range ingredients[i] {
            if _, exists := graph[ingredient]; !exists {
                graph[ingredient] = make([]string, 0)
            }
            graph[ingredient] = append(graph[ingredient], recipes[i])
        }
    }
    queue := make([]string, len(supplies))
    for i := range supplies {
        queue[i] = supplies[i]
    }
    for nodeKey := range inDegree {
        if inDegree[nodeKey] == 0 {
            queue = append(queue, nodeKey)
        }
    }
    res := make([]string, 0, len(recipes))
    var node string
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            if _, exists := recipeSet[node]; exists {
                res = append(res, node)
                delete(recipeSet, node)
            }
            // decrease indegree of recipe dst node
            neighbors := graph[node]
            if len(neighbors) != 0 {
                for _, nei := range neighbors {
                    inDegree[nei]--
                    if inDegree[nei] == 0 {
                        queue = append(queue, nei)
                    }
                }
            }
        }
    }
    
    return res
}