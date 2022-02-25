
// 279. Perfect Squares


// bfs approach
// queue len determines the space complexity here
// maxWidth occurs at final lvl thus, we allocate up to O(n)
// where n is the number of nodes in our n-ary tree
// time: O(n) where n is numNodes


func numSquares(n int) int {
    perfectSquares := make([]int, 0, n >> 1)
    num := 1
    // generate perfect squares
    for num * num <= n {
        perfectSquares = append(perfectSquares, num*num)
        num++
    }      
    visited := make([]bool, n+1)
    var (
        queue []int
        node  int
    )
    queue = append(queue, n)
    visited[n] = true
    leastNumSquares := 0
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            if node == 0 {
                return leastNumSquares
            }
            for i := len(perfectSquares)-1; i >= 0; i-- {
                if diff := node - perfectSquares[i]; diff >= 0 &&
                !visited[diff] {
                    visited[diff] = true
                    queue = append(queue, diff)
                }
            }
        }
        leastNumSquares++
    }
    return -1
}