//time: O(n)
//space: O(1)

//preconditions: 
//nums is sorted in non-decreasing order
//input test cases generate exactly 1 solution 
//elem can't be used twice

func twoSum(numbers []int, target int) []int {
    res := make([]int, 2)
    n := len(numbers)
    left, right := 0, n-1 
    for left != right {
        sumVals := numbers[left] + numbers[right]
        switch {
        case sumVals < target:
            left++
        case sumVals > target:
            right--
        default: //==target
            //return indices as 1-indexed values
            res[0] = left+1
            res[1] = right+1
            return res
        }
    }
    return res
}