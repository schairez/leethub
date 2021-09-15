/*
time: O(len(nums1) + len(nums2)) ~ O(n)
space: O(len(nums2)) ~ O(n)
*/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    stack := []int{}
    res := make([]int, len(nums1))
    //map w/ keys as num and next greatest num as val
    nextGreatest := make(map[int]int)
    for _, val := range nums2 {
        for len(stack) > 0 && val > stack[len(stack)-1] {
            lifoElem := stack[len(stack)-1]
            nextGreatest[lifoElem] = val
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, val)
    }
    for i := range res {
        nextG, ok := nextGreatest[nums1[i]] 
        if !ok {
            res[i] = -1
            continue
        }
        res[i] = nextG
    }
    return res
    
}
