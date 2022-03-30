// 1726. Tuple with Same Product
// time: O(n^2)
// space: O(n)

func tupleSameProduct(nums []int) int {
    n := len(nums)
    res := 0
    productMap := make(map[int]int)
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            key := nums[i]*nums[j]
            if v, ok := productMap[key]; ok {
                res += v
            }
            productMap[key]++
        }
    } 
    // 2! * 2! ways for each pairing 
    return 8 * res
}