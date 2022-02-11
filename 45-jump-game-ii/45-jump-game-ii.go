// explore based on number of steps taken
// shortest path approach is a marker for BFS


func min(a, b int) int { if a <= b { return a}; return b}


//BFS w/o queue approach
//time: O(n) 
//space: O(n)

func jump(nums []int) int {
    n := len(nums)
    if n == 1 {
        return 0
    }
    distToNode := make([]int, n) //dp array
    lastIdx := len(nums) - 1
    visited := make([]bool, n)
    nextIdxToJump := 0
    visited[nextIdxToJump] = true
    distToNode[nextIdxToJump] = 0
    for idx, val := range nums {
        nextIdxToJump = min(val + idx, lastIdx)
        for nextIdxToJump != idx {
            if !visited[nextIdxToJump] {
                if nextIdxToJump == lastIdx {
                    return distToNode[idx] + 1
                }
                distToNode[nextIdxToJump] = distToNode[idx] + 1
                visited[nextIdxToJump] = true
            } 
            nextIdxToJump--
        }
    }
    return distToNode[lastIdx]
}


/*
//BFS with queue approach
//time: O(n) 
//space: O(n)
func jump(nums []int) int {
    n := len(nums)
    if n == 1 {
        return 0
    }
    lastIdx := len(nums) - 1
    currNode, numSteps := 0, 0
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
*/