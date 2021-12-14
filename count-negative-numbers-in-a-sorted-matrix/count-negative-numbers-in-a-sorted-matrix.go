/*
 
[[4, 3, 2,-1]]
[[4, 3, 2,-1]]
 [3, 2, 1,-1]
 [1, 1, -1,-2]
 
 [-1,-1,-2,-3]]
         <----
 [3, 2, 1,-1]
        V
 [1, 1, -1,-2]
  <------V
 [-1,-1,-2,-3]]
*/

//search space reduction v
//time : O(m + n)
//space: O(1)
func countNegatives(grid [][]int) int {
    numR := len(grid)
    numC := len(grid[0])
    res := 0
    col := len(grid[0])-1
    for row := 0; row < numR; row++ {
        for col >= 0 && grid[row][col] < 0 {
            col--
        }
        res += numC-1 - col
    }
    return res
}


//[4,3,2,-1,-2]
//hi = 5, lo = 0; 
//mid = 5//2 = 2
//arr[mid] >=0; lo = mid+1; lo = 3
//mid = 3 + (5-3)//2 = 4;
//arr[mid] = -2 < 0; hi = mid = 4
//mid = 3 + (4-3)//2 = 3; arr[mid] < 0; hi = 3
//leftMost neg @ idx 3
//numNeg at row = 5 - 3 = 2 
//binSearch approach
//time:  O(nlog(m))
//space: O(1)
func countNegativesBinSearch(grid [][]int) int {
    numR, numC := len(grid), len(grid[0])
    
    //returns num neg vals at curR
    //time: O(log(numC)) for given row input
    numNegValsAtRow := func(curR int) int {
        lo, hi := 0, numC
        //no negative at row
        if grid[curR][hi-1] >= 0 {
            return 0
        }
        //all negative at row
        if grid[curR][lo] < 0 {
            return numC
        }
        //start assuming entire r is neg
        //find leftmost negative number
        for lo != hi  {
            mid := lo + (hi-lo) >> 1
            if grid[curR][mid] >= 0 {
                lo = mid+1
            } else {
                hi = mid
            } 
        }
        numNegative := numC - lo
        return numNegative
    }
    res := 0
    curR := 0
    for curR < numR {
        res += numNegValsAtRow(curR)
        curR++
    } 
    return res
}
