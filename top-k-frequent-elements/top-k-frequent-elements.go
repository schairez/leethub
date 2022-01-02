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
    for numKey, cntVal := range freqMap {
        if _, ok := invFreq[cntVal]; !ok {
            invFreq[cntVal] = []int{}
        } 
        invFreq[cntVal] = append(invFreq[cntVal], numKey)
    }
    res := make([]int, 0, k)
    for currFreq := maxFreq; currFreq >=0; currFreq-- {
        if len(res) == k {
            break
        }
        if freqElemsArr, ok := invFreq[currFreq]; ok {
            res = append(res, freqElemsArr...)
        }
    }
    
    
    
    return res
    
}

