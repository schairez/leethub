// 1 < = len(nums) <= 20
// 2^20
//
//
//

func findTargetSumWays(nums []int, target int) int {
    n := len(nums)
   // var nums [1 << 20]int
    var (
        //queue []int
        node  int
        lvl   int
        res   int
    )
    queue := make([]int, 0, 1 << 20)
    queue = append(queue, []int{nums[0], -nums[0]}...)
    for len(queue) != 0 {
        lvl++
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            if lvl == n {
                if node == target {
                    res++
                }
                continue
            }
            queue = append(queue,
                           []int{node+nums[lvl],node-nums[lvl]}...)
        }
    }
    return res
}


