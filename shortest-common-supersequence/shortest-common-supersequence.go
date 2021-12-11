import "strings"
/*
lcs time:  O(n*m)
lcs space: O(n*m)

deriving superstring time: O(n+m)
deriving superstring space: O(n+m)
after reversing onto other str buffer O(n+m) time and space

therefore,
time: O(n*m)
space: O(n*m)

*/


func max(a, b int) int { if a >= b { return a}; return b }

func shortestCommonSupersequence(str1 string, str2 string) string {
    //get lcs
    lenStr1, lenStr2 := len(str1), len(str2)
    dp := make([][]int, lenStr1+1)
    for row := range dp {
        dp[row] = make([]int, lenStr2+1)
        dp[row][0] = 0
    }
    for col := 0; col < len(dp[0]); col++ {
        dp[0][col] = 0
    }
    for row := 1; row <= lenStr1; row++ {
        for col :=1; col <= lenStr2; col++ {
            switch {
            case str1[row-1] == str2[col-1]:
                dp[row][col] = dp[row-1][col-1] + 1
            default:
                fromLeft := dp[row][col-1]
                fromAbove := dp[row-1][col]
                dp[row][col] = max(fromLeft, fromAbove)
            }
        }
    }
    //append vals to string buffer
    var sb strings.Builder
    currIdxStr1, currIdxStr2 := lenStr1-1, lenStr2-1
    for currIdxStr1 != -1 || currIdxStr2 != -1 {
        switch {
        case (currIdxStr1 == -1) != (currIdxStr2 == -1):
            if currIdxStr1 < 0 {
                sb.WriteByte(str2[currIdxStr2])
                currIdxStr2--
            } else {
                sb.WriteByte(str1[currIdxStr1])
                currIdxStr1--
            }
        case str1[currIdxStr1] == str2[currIdxStr2]:
            sb.WriteByte(str1[currIdxStr1])
            currIdxStr1--
            currIdxStr2--
        default:
            if dp[currIdxStr1+1][currIdxStr2] >= 
                dp[currIdxStr1][currIdxStr2+1] {
                sb.WriteByte(str2[currIdxStr2])
                currIdxStr2--
            } else {
                sb.WriteByte(str1[currIdxStr1])
                currIdxStr1--
            }
        }
    }
    
    reversedVersion := sb.String()
    
    var sbOrdered strings.Builder
    sbOrdered.Grow(len(reversedVersion))
    for i := len(reversedVersion) -1; i >= 0; i -- {
        sbOrdered.WriteByte(reversedVersion[i])
    }
    
    return sbOrdered.String()
    
}
