//time: O(n)
//space: O(1)

func max(a, b int) int { if a >= b { return a}; return b}

func findMaxAverage(nums []int, k int) float64 {
    n := len(nums)
    var maxSum, localSum int
    for i := 0; i < k; i++ {
        localSum += nums[i]
    }
    maxSum = localSum
    for i := k; i < n; i++ {
        localSum -= nums[i-k]
        localSum += nums[i]
        maxSum = max(maxSum, localSum)
    }
    return float64(maxSum) / float64(k)
}







































//sliding window approach
//time: O(n)
//space: O(n)

/*
       st       curr
                 |
nums = [1,12,-5,-6,50,3], k = 4

*/

/*
func findMaxAverage(nums []int, k int) float64 {
    avgs := []float64{}
    var avg float64
    start, sum := 0,0
    for currIdx, val := range nums {
        sum += val
        if currIdx - start + 1 == k {
            avg = float64(sum)/ float64(k) 
            avgs = append(avgs, avg)
            sum -= nums[start]
            start+=1
        } 
    }
    return max(avgs...)
    
}


func max(items ...float64) float64 {
    maxSoFar := items[0]
    for i:=1; i < len(items); i++ {
        if items[i] > maxSoFar {
            maxSoFar = items[i]
        }
    }
    return maxSoFar
}
*/