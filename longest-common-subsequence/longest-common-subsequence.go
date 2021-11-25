//time: O(m*n); m=text1len; n=text2len
//space: O(m*n)


//state transition fn
// either incr when common chars, else the max b/w considering prev commonSubseq
// with i-1 char in text1 or the j-1 char in text 2 
// dp[i][j] = { if text1[i] == text2[j] then -> 1 + dp[i-1][j-1]
//              else max(dp[i-1][j], dp[i][j-1]) 
//            }

func max(a, b int) int { if a>=b { return a }; return b } 

func longestCommonSubsequence(text1 string, text2 string) int {
    text1Len, text2Len := len(text1), len(text2)
    //[i,0] = 0; the case of considerign an empty substr
    dp := make([][]int, text1Len+1)
    for row := range dp {
        dp[row] = make([]int, text2Len+1)
        dp[row][0] = 0
    }
    for row:=1; row <= text1Len; row++ {
        for col:=1; col <= text2Len; col++ {
            switch {
            case text1[row-1] == text2[col-1]:
                dp[row][col] = 1 + dp[row-1][col-1]
            default:
                dp[row][col] = max(dp[row-1][col], dp[row][col-1])
            }
        }
    }
    return dp[text1Len][text2Len]
}