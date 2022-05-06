func canCross(stones []int) bool {
    dirs := [3]int{-1, 0, 1}
    n := len(stones)
    src := stones[0]
    dst := stones[n-1]
    stonesMap := make(map[int]struct{}, n)
    for i := range stones {
        stonesMap[stones[i]] = struct{}{}
    }
    type nodeData struct {loc, k int}
    var (
        pollNode nodeData 
        queue []nodeData
    )
    visited := make(map[nodeData]bool, n)
    visited[nodeData{src,0}] = true
    queue = append(queue, nodeData{src, 0})
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            pollNode, queue = queue[0], queue[1:]
            if pollNode.loc == dst {
                return true
            }
            for i := 0; i < 3; i++ {
                jump := pollNode.k + dirs[i]
                nextLoc := pollNode.loc + jump
                data := nodeData{nextLoc, jump}
                if _, exists := stonesMap[nextLoc]; exists {
                    if !visited[data] {
                        visited[data] = true
                        queue = append(queue, data)
                    }
                }
            }
        }
    }
    return false
}