//m = numRows; n = numCols
//space: O(1)
//time: O(mlog(n))

func searchMatrix(matrix [][]int, target int) bool {
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