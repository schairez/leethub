/*
time: O(n)
space: O(n)
*/


func max(a, b int) int { if a >=b { return a}; return b}

func topKFrequent(nums []int, k int) []int {
    if len(nums) == 1 {
        return nums
    }
    maxFreq := 0
    //freqMap of num to cnt
    freqMap := make(map[int]int)
    for _, num := range nums {
        if _, ok := freqMap[num]; !ok {
            freqMap[num] = 0 
        }
        freqMap[num]++
        maxFreq = max(maxFreq, freqMap[num])
    }
    //inverse frequency
    //maps freq cnt's to arrlist of nums
    invFreq := map[int][]int{}
    for numKey, freqVal := range freqMap {
        if _, ok := invFreq[freqVal]; !ok {
            invFreq[freqVal] = []int{}
        } 
        invFreq[freqVal] = append(invFreq[freqVal], numKey)
    }
    res := make([]int, 0, k)
    for currFreq := maxFreq; currFreq >=0; currFreq-- {
        if freqElemsArr, ok := invFreq[currFreq]; ok {
            for i := len(freqElemsArr)-1; i >= 0 && k > 0; i-- {
                res = append(res, freqElemsArr[i])
                k--
            }
            if k == 0 { break }
        }
    }
    
    
    
    return res
    
}

