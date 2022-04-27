
func max(a, b int) int { if a >= b { return a}; return b}

// 45. Jump Game II
// time: O(n) space: O(1)
// condition: we can reach lastIdx
func jump(nums []int) int {
    n := len(nums)
    dst := n-1
    numJumps := 0
    currLvlEnd, furthestIdx := 0, 0
    for i := 0; i < n-1; i++ {
        furthestIdx = max(furthestIdx, i + nums[i])
        if i == currLvlEnd {
            numJumps++
            if furthestIdx >= dst {
                return numJumps
            }
            currLvlEnd = furthestIdx
        }
    }
    return numJumps
}




func jumpBFS(nums []int) int {
    lvl := 0
    n := len(nums)
    var (
        node  int
        queue []int
    )
    visited := make([]bool, n)
    queue = append(queue, 0)
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            if node == n-1 {
                return lvl
            }
            furthestIdx := min(node + nums[node], n-1)
            for nextIdx := 1 + node; nextIdx <= furthestIdx; nextIdx++ {
                if !visited[nextIdx] {
                    queue = append(queue, nextIdx)
                    visited[nextIdx] = true
                }
            }
        }
        lvl++
    }
    return -1
}

func min(a, b int) int { if a <= b { return a}; return b}


// top down dfs memo approach
// time: O(n^2) space: O(n)
func jumpDFS(nums []int) int {
    n := len(nums)
    memo := make([]int, n)
    for i := range memo {
        memo[i] = 10_000
    }
    var dfs func(int) int
    dfs = func(idx int) int {
        if idx >= n - 1 {
            return 0
        }
        if memo[idx] != 10_000 {
            return memo[idx]
        }
        minJumps := 10_000
        nextIdx, furthestIdx := idx + 1, idx + nums[idx]
        for nextIdx <= furthestIdx {
            minJumps = min(minJumps, 1 + dfs(nextIdx)) 
            nextIdx++
        }
        memo[idx] = minJumps
        return memo[idx]
    }
    return dfs(0)
}