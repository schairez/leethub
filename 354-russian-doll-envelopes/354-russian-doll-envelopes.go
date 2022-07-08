
func max(a, b int) int { if a >= b { return a}; return b}

func isEnv2Larger(env1 [2]int, env2 []int) bool {
    return env2[0] > env1[0] && env2[1] > env1[1]
}

func maxEnvelopes(envelopes [][]int) int {
    n := len(envelopes)
    sort.Slice(envelopes, func(i, j int) bool {
        if envelopes[i][0] == envelopes[j][0] {
            return envelopes[i][1] > envelopes[j][1]
        }
        return envelopes[i][0] < envelopes[j][0]
    })
    dp := make([][2]int, n+1)
    k := 0
    const minInt32 = -1 << 31
    dp[k][0], dp[k][1] = minInt32, minInt32
    for i := 0; i < n; i++ {
        if isEnv2Larger(dp[k], envelopes[i]) {
            k++
            dp[k][0], dp[k][1] = envelopes[i][0], envelopes[i][1]
        } else {
            lo, hi := 0, k
            for lo + 1 < hi {
                mid := lo + (hi-lo) >> 1
                if !isEnv2Larger(dp[mid], envelopes[i]) {
                    hi = mid
                } else {
                    lo = mid
                }
            }
            dp[hi][0], dp[hi][1] = envelopes[i][0], envelopes[i][1]
        }
    }
    return k
}