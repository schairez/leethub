// space: O(N)
// time: O(N)
func findDiagonalOrder(nums [][]int) []int {
    numR := len(nums)
    if numR == 1 {
        return nums[0]
    }
    res := []int{}
    //curR, curC pair
    q := [][2]int{}
    q = append(q, [2]int{0,0})
    for len(q) > 0 {
        curV := q[0]
        q = q[1:]
        res = append(res, nums[curV[0]][curV[1]])
        //add neighbors to q
        //nei down if in col0
        if nextR := curV[0]+1; nextR <= numR-1 && curV[1] == 0 {
            q = append(q, [2]int{nextR, curV[1]})
        }
        //nei right
        if nextC := curV[1]+1; nextC < len(nums[curV[0]]) {
            q = append(q, [2]int{curV[0], nextC})
        }
    }
    return res
} 
    
    
    
    
    
    
    
    
    
    
/*    
//prev version; passes 53/56 Test Cases
//1 <= nums[i][j] <= 10^9
func findDiagonalOrder(nums [][]int) []int {
    if len(nums) == 1 {
        return nums[0]
    }
    //if non-sparse m*n
    numR := len(nums)
    numC := 0
    //check if column-dense, row-sparse
    //find max column width
    for i := range nums {
        curWidth := len(nums[i])
        if curWidth > numC {
            numC = curWidth
        }
    }
    res := []int{}
    curR, curC := 0, 0
    isValid := func(curR, curC int) bool {
        if curR < 0 || curC > len(nums[curR])-1 {
            return false
        }
        return true
    }
    rightDiagonalVals := func(curR, curC int)  {
        //[2][0]
        //[1][1]
        nextR, nextC := curR, curC
        for nextR >= 0 {
            nextR--
            nextC++
            if ok := isValid(nextR, nextC); !ok {
                continue
            }
            res = append(res, nums[nextR][nextC])
        }
    }
    //does row contain vals
    
    //number of parallel lines
    totalDiagonals := numR + numC - 1
    numDiagonals := 0
    moveDownRows := true
    for numDiagonals < totalDiagonals {
        switch {
        case moveDownRows:
            if ok := isValid(curR, curC); ok {
                res = append(res, nums[curR][curC]) 
            } 
            rightDiagonalVals(curR, curC)
            curR++
            numDiagonals++
            if curR == numR { 
                moveDownRows = false 
                curR--
                curC++
            }
        default:
            if ok := isValid(curR, curC); ok {
                res = append(res, nums[curR][curC]) 
            } 
            rightDiagonalVals(curR, curC)
            curC++
            numDiagonals++
        }
    } 
    return res
}
*/