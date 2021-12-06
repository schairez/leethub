

func min(vals ...int) int { 
    minV := vals[0]
    for i := range vals {
        if vals[i] < minV {
            minV = vals[i]
        }
    }
    return minV
}


//2D DP table bottom up
//time: O(len(w1)*len(w2))
//space: O(len(w1)*len(w2))
func minDistance(word1 string, word2 string) int {
    lenW1, lenW2 := len(word1), len(word2)
    dp := make([][]int, lenW1+1)
    for row := range dp {
        dp[row] = make([]int, lenW2+1)
        dp[row][0] = row
    }
    for col := 0; col < lenW2+1; col++ {
        dp[0][col] = col
    }
    for row := 1; row < lenW1+1; row++ {
        for col := 1; col < lenW2+1; col++ {
            if word1[row-1] == word2[col-1] {
                dp[row][col] = dp[row-1][col-1]
            } else { //delete, insert, replace ops
                insertOper := dp[row-1][col]
                replaceOper := dp[row-1][col-1]
                deleteOper := dp[row][col-1]
                dp[row][col] = min(insertOper, replaceOper,
                                  deleteOper) + 1
            }
        }
    }
    return dp[lenW1][lenW2]
}



//top-down memoized DP
//time: O(len(w1)*len(w2))
//space: O(len(w1)*len(w2))
func minDistanceMemo(word1 string, word2 string) int {
    genKey := func(i, j int) string { return string(i)+"|"+string(j) }
    
    memo := make(map[string]int)
    
    var editDistance func(int, int) int
    editDistance = func(idxW1, idxW2 int) int {
        if idxW1 == 0 {
            return idxW2
        } else if idxW2 == 0 {
            return idxW1
        }
        key := genKey(idxW1, idxW2)
        if minOper, ok := memo[key]; ok {
            return minOper
        }
        //no operation necessary
        if word1[idxW1-1] == word2[idxW2-1] {
            memo[key] = editDistance(idxW1-1, idxW2-1)
        } else { //replace, delete, insert operations
            replaceOper := editDistance(idxW1-1, idxW2-1)
            insertOper := editDistance(idxW1, idxW2-1)
            deleteOper := editDistance(idxW1-1, idxW2)
            memo[key] = min(replaceOper, insertOper, deleteOper) + 1
            
        }
        return memo[key]
    }
    
    lenW1, lenW2 := len(word1), len(word2)
    return editDistance(lenW1, lenW2)
    
}

/*

Given two strings word1 and word2, return the minimum number
of operations required to convert word1 to word2.

You have the following three operations permitted on a word:
Insert a character
Delete a character
Replace a character

ex1
Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation: 
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')

*/

