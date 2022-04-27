// 1306. Jump Game III
// time: O(n)
// space: O(n)
func canReach(arr []int, start int) bool {
    n := len(arr)
    visited := make([]bool, n)
    var (
        node  int
        queue []int
    )
    visited[start] = true
    queue = append(queue, start)
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            if arr[node] == 0 {
                return true
            }
            // at index i, you can jump to i + arr[i] or i - arr[i
            if nei := node + arr[node]; nei < n && !visited[nei] {
                queue = append(queue, nei)
                visited[nei] = true
            }
            if nei := node - arr[node]; nei >= 0 && !visited[nei] {
                queue = append(queue, nei)
                visited[nei] = true
            }
        }
    }
    return false
}