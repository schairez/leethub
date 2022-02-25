
func min(a, b int) int { if a <= b { return a}; return b}

func coinChange(coins []int, amount int) int {
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

















































/*
//Return the fewest number of
coins that you need to make up that
amount. If that amount of money cannot be made
up by any combination of the coins, return -1.

//goal: minimize number of coins s.t. dollarValue == amount

Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1

constraints:
1 <= coins.length <= 12
1 <= coins[i] <= 2^31 -1
0 <= amount <= 10^4

*/

//state transition
//dp[i] = min(dp[i], dp[i - coins[i]] + 1)

//time: O(amount*len(coins))
//space: O(amount)
/*

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