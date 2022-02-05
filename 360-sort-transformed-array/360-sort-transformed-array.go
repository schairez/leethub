
//360. Sort Transformed Array

//f(x) = ax^2 + bx + c
func quadWithCoeff(a, b, c int) func(x int) int {
    return func(x int) int {
        return a * (x * x) + (b * x) + c
    }
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

//time: O(n)
//space: O(1) if not considering res; O(n) if so

//precondition: nums is sorted in non-decr order
func sortTransformedArray(nums []int, a int, b int, c int) []int {
    n := len(nums)
    res := make([]int, n)
    //quad with coefficients initialized
    quadFn := quadWithCoeff(a, b, c)
    idx := n-1
    //upside down parabola
    //if upside down parabola; tails are min values
    //else: tails are max values
    invParabolaCond := a < 0
    if invParabolaCond {
        idx = 0
    }
    var quadLeftV, quadRightV int
    left, right := 0, n-1
    for left <= right {
        quadLeftV, quadRightV = quadFn(nums[left]), quadFn(nums[right])
        if invParabolaCond {
            if quadLeftV < quadRightV {
                res[idx] = quadLeftV
                left++
            } else {
                res[idx] = quadRightV
                right--
            }
            idx++
        } else {
            if quadLeftV > quadRightV {
                res[idx] = quadLeftV
                left++
            } else {
                res[idx] = quadRightV
                right--
            }
            idx--
        }
    }
    return res
}
