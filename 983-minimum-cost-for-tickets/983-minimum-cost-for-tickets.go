
func max(a, b int) int { if a >= b { return a}; return b}
func min(a, b int) int { if a <= b { return a}; return b}

func mincostTickets(days []int, costs []int) int {
    n := len(days)
    travelDaySet := make(map[int]struct{}, n)
    for i := range days {
        travelDaySet[days[i]] = struct{}{}
    }
    var dp [366]int 
    // 1 <= days[i] <= 365
    dp[0] = 0
    for day := 1; day <= 365; day++ {
        if _, exists := travelDaySet[day]; !exists {
            dp[day] = dp[day-1]
        } else {
            day1Cand := dp[day-1] + costs[0] 
            day7Cand := dp[max(day-7, 0)] + costs[1] 
            day30Cand := dp[max(day-30, 0)] + costs[2] 
            dp[day] = min(min(day1Cand, day7Cand), day30Cand)
        }
    }
    return dp[365]
    
}