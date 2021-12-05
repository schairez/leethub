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
func min(a, b int) int { if a <= b { return a }; return b }

//state transition
//dp[i] = min(dp[i], dp[i - coins[i]] + 1)

//time: O(amount*len(coins))
//space: O(amount)
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
