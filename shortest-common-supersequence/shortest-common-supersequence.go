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




/*
str1 = "abac", str2 = "cab"
Output: "cabac"

iterStr1 rep by i
iterStr2 rep by j
1. find lcs
//lcs = {"a","b"}
2. zip iter through both strings
"c"
"ca"
"cab"
"caba"
"cabac"
3. rev res
"cabac"

str1 = "geek"
str2 = "eke"
lcs = "ek"
res = "geeke"
-------------------
"e"
"ek"
"eke"
"ekee"
"ekeeg"
after rev
"geeke"
lcs tabulation
recurrence formula 
f(i,j) = {
          f(i-1, j-1) + 1 if str1[i] == str2[j]}
              //fromAbove, fromLeft
          max(f(i-1,j), f(i, j-1))

     0  g e e k
   0 0  0 0 0 0 
   e 0  0 1 1 1 
   k 0  0 1 1 2
   e 0  0 1 2 2


i = m-1
j = n-1
for i != -1 || j != -1 {
    //t != f || f != t
    if (i < 0) != (j < 0) {
        if i < 0 { res += str2[j]; j--}
        else { res+=str1[i]; i--}
    }
    elif str1[i] == str2[j] {
        res += str1[i]; i--; j--
    }
    else {
        if dp[i+1][j] >= dp[i][j+1] { res+=str2[j]; j--}
        else { res+=str1[i]; i--}
        }
}

    
*/





