/*
time: O(n)
space: O(n)
*/

import "fmt"

func topKFrequent(nums []int, k int) []int {
    if len(nums) == 1 {
        return nums
    }
    res := make([]int, 0, k)
    freqMap := make(map[int]int)
    for _, num := range nums {
        if _, ok := freqMap[num]; !ok {
            freqMap[num] = 0 
        }
        freqMap[num]++
    }
    //inverse frequency
    invFreq := map[int][]int{}
    for k,v := range freqMap {
        if _, ok := invFreq[v]; !ok {
            invFreq[v] = []int{}
        } 
        invFreq[v] = append(invFreq[v], k)
    }
    fmt.Println(invFreq)
    for i:=len(nums); i >=0; i-- {
        if len(res) == k {
            break
        }
        if freqElemsArr, ok := invFreq[i]; ok {
            res = append(res, freqElemsArr...)
        }
    }
    
    
    
    return res
    
}

