//search space reduction v
//time : O(m + n)
//space: O(1)
func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])
    curR, curC := m-1, 0
    for curR >= 0 && curC < n {
        curV := matrix[curR][curC] 
        switch {
        case curV == target:
            return true 
        case curV > target:
            curR--
        case curV < target:
            curC++
        }
    }
    return false
}


//m = numRows; n = numCols
//space: O(1)
//time: O(mlog(n))
func searchMatrixBinSearchV(matrix [][]int, target int) bool {
    binSearch := func(arr []int, t int) bool {
        lo, hi := 0, len(arr) -1
        for lo <= hi {
            m := lo + (hi - lo) >> 1 
            if arr[m] == t { 
                return true 
            } else if arr[m] > t {
                hi = m -1
            } else {
                lo = m +1
            }
        }
        return false
    }
    m := len(matrix)
    for y:=0; y < m; y++ {
        if found := binSearch(matrix[y], target); found {
            return true
        } 
    }
    return false
}