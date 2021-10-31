import "math"

//kadene's algo inspired
//time:O(N)
//space:O(1)

func maxProduct(nums []int) int {
    minProduct, maxProduct := 1,1
    res := math.MinInt32 
    for i:=1;i<=len(nums);i++ {
        tmp := minProduct
        minProduct = min(nums[i-1], min(nums[i-1]*minProduct, nums[i-1]*maxProduct))
        maxProduct = max(nums[i-1], max(nums[i-1]*maxProduct, nums[i-1]*tmp))
        res = max(res, maxProduct)
        
    }
    
    return res
    
}

func min(a,b int) int {
    if a <= b { 
        return a
    }
    return b
}
                         
                         

func max(a, b int) int {
    if a >= b {
        return a
    }
    return b
}