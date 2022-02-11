// explore based on number of steps taken
// shortest path approach is a marker for BFS
//
//
//
//

func min(a, b int) int { if a <= b { return a}; return b}

func jump(nums []int) int {
    n := len(nums)
    if n == 1 {
        return 0
    }
    lastIdx := len(nums) - 1
    numSteps := 0
    currNode := 0
    visited := make([]bool, n)
    visited[0] = true
    //enqueue startNode @ idx=0
    queue := []int{0}
    for len(queue) != 0 {
        numNodesAtLvl := len(queue)
        for numNodesAtLvl != 0 {
            currNode, queue = queue[0], queue[1:]
            nextNumJumps := min(nums[currNode], lastIdx-currNode)
            for nextNumJumps != 0 {
                nextLvlNode := currNode + nextNumJumps
                if nextLvlNode == lastIdx {
                    return numSteps+1
                }
                if !visited[nextLvlNode] {
                    visited[nextLvlNode] = true
                    queue = append(queue, nextLvlNode)
                }
                nextNumJumps--
            }
            numNodesAtLvl--
        }
        numSteps++
    }
    return numSteps+1
}