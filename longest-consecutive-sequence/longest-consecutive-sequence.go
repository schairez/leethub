//time: O(n)
//space: O(n); hashMap


//Input: nums = [100,4,200,1,3,2]
//

func longestConsecutive(nums []int) int {
    max := func(a, b int) int { 
        if a >= b {
            return a
        }
        return b
    }
    globalLongest := 0
    numSet := make(map[int]bool)
    for _, num := range nums {
        numSet[num] = true
    }
    for num := range numSet {
        if _, ok := numSet[num-1]; !ok {
            currNum := num
            localLongest := 1
            nextNum := currNum+1
            currInSetCond := numSet[nextNum] 
            for currInSetCond {
                nextNum++
                localLongest++
                currInSetCond = numSet[nextNum]
            }
            globalLongest = max(localLongest, globalLongest)
        } 
    } 
    return globalLongest
}