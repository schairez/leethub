
func min(a, b int) int { if a <= b { return a}; return b}

//BFS approach
// time: O(len(coins) * |amount|)
// space: O(n) for queue and visited set

func coinChange(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }
    visited := make([]bool, amount+1)
    visited[amount] = true
    var (
        queue []int
        nodeAmtBalance  int
        numCoins int
    )
    queue = append(queue, amount)
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            nodeAmtBalance, queue = queue[0], queue[1:]
            if nodeAmtBalance == 0 {
                return numCoins
            }
            for _, denomVal := range coins {
                if remBalance := nodeAmtBalance - denomVal; remBalance >= 0 &&
                !visited[remBalance] {
                    visited[remBalance] = true
                    queue = append(queue, remBalance)
                }
            }
        }
        numCoins++
    }
    return -1
}
 

// dp approach
// space: O(n) space for dp slice
// time: O(|amount| * |coins|) 

func coinChangeDP(coins []int, amount int) int {
    const maxInt32 = 1 << 31 -1
    dp := make([]int, amount+1)
    for amt := 1; amt < len(dp); amt++ {
        dp[amt] = maxInt32
    }
    for _, denomination := range coins {
        for currAmt := 1; currAmt <= amount; currAmt++ {
            if remAmt := currAmt - denomination; remAmt >= 0 {
                dp[currAmt] = min(dp[currAmt], 1 + dp[remAmt])
            }
        }
    } 
    if dp[amount] != maxInt32 {
        return dp[amount]
    }
    return -1
}



//state transition
//dp[i] = min(dp[i], dp[i - coins[i]] + 1)

//time: O(amount*len(coins))
//space: O(amount)
/*
//version from 12/5/2021

func min(a, b int) int { if a <= b { return a }; return b }
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    //we set each idx for our coinAmnt data as 1 + input amt
    //a value out of our input scope
    maxVal := amount+1
    for coinAmtIdx := range dp {
        dp[coinAmtIdx] = maxVal
    }
    dp[0] = 0 //0 coins needed for $0
    for moneyAmt := 1; moneyAmt <= amount; moneyAmt++ {
        for denomIdx := range coins {
            if coinVal := coins[denomIdx]; coinVal <= moneyAmt {
                dp[moneyAmt] = min(dp[moneyAmt], dp[moneyAmt - coinVal] + 1)
            }
        }
    }
    if res := dp[amount]; res != maxVal {
        return res
    }
    return -1
}
*/