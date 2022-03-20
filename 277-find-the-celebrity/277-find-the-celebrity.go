/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
// 0 to n-1
// [[1,0],[0,1]]

func solution(knows func(a int, b int) bool) func(n int) int {
    return func(n int) int {
        cand := 0
        for i := 0; i < n; i++ {
            if cand == i {
                continue
            }
            if knows(cand, i) {
                cand = i
            }
        }
        for i := 0; i < n; i++ {
            if cand == i {
                continue
            }
            if knows(cand, i) {
                return -1
            }
            if !knows(i, cand) {
                return -1
            }
        }
        return cand
    }
}