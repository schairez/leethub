//You are given an integer array coins representing coins
//of different denominations and an integer amount representing a total amount of money.
//Return the number of combinations that make up that amount.
//If that amount of money cannot be made up by any combination of the coins, return 0.

//time: O(amount*len(coins))
//space: O(amount)

func change(amount int, coins []int) int {
    dp := make([]int, amount+1)
    //1 possible combination of $0
    dp[0] = 1
    for coinIdx := range coins {
        coinVal := coins[coinIdx]
        moneyAmt := coinVal
        for ; moneyAmt <= amount; moneyAmt++ {
            dp[moneyAmt] += dp[moneyAmt - coinVal]
        }
    }  
    return dp[amount]
    
}


/*
Input: amount = 5, coins = [1,2,5]
Output: 4
Explanation: there are four ways to make up the amount:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
iter coin amt idx's i 
  iter coinVal to amnt j
  dp[j] += dp[j - amt[i]]

|0||1||2||3||4||5|
------------------
|1||1||1||1||1||1| ; after coinVal ->1
|1||1||2||2||3||3| ; after coinVal ->2
|1||1||2||2||3||4| ; after coinVal ->5

options for $4 after coinVal of $2 iteration -> {1,1,1,1}, {2,1,1}, {2,2}
options for $5 after coinVal of $2 iteration -> {1,1,1,1,1}, {2,1,1,1}, {2,2,1}
options for $5 after coinVal of $5 iteration -> {1,1,1,1,1}, {2,1,1,1}, {2,2,1}, {5}

*/



