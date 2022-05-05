/*
nums1 = [4,1,2], nums2 = [1,3,4,2]
Output: [-1,3,-1]
*/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
    n1 := len(nums1)
    n2 := len(nums2)
    m := make(map[int]int)
    var pollNode int
    var stack []int
    /*
       st = [0]   @i=1 -->  m[1] = 3; st = [1] ; @i=2 m[3] = 4, st=[2]
       @i=3  st = [2,3]
    */
    for i := 0; i < n2; i++ {
        for len(stack) != 0 && nums2[stack[len(stack)-1]] < nums2[i] {
            pollNode, stack = stack[len(stack)-1], stack[:len(stack)-1]
            // m[1] = 3
            m[nums2[pollNode]] = nums2[i]
        }
        stack = append(stack, i)
    }
    res := make([]int, n1)
    for i := 0; i < n1; i++ {
        if nextGreater, exists := m[nums1[i]]; exists {
            res[i] = nextGreater
            continue
        }
        res[i] = -1
    }
    return res
}