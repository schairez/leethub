import "sort"

//k = freq of unique elems
//time: O(nlogn + nlogk) ~ O(nlogn); 
//space: O(n)
/*
testCases
[6,3,0,1]
[0,0,0,0]

*/
func findOriginalArray(changed []int) []int {
    n := len(changed)
    if n % 2 != 0 {
        return []int{}
    }
    sort.Ints(changed) //O(nlogn)
    expSizeRes := n /2
    res := make([]int, 0, expSizeRes)
    valFreqMap := make(map[int]int, n)
    for _, val := range changed {
        valFreqMap[val]++
    }
    for len(valFreqMap) != 0 {
        for _, val := range changed {
            cnt := valFreqMap[val]
            switch {
            case val == 0:
                if cnt % 2 != 0 { //odd num vals
                    return []int{}
                }
                for i:= 0; i < cnt/2; i++ {
                    res = append(res, val)
                }
                delete(valFreqMap, val)
            default:
                //ex: [3,3,3,6,6,6]
                if valFreqMap[val*2] < cnt {
                    return []int{}
                }
                delete(valFreqMap, val)
                for i := 0; i < cnt; i++ {
                    res = append(res, val)
                }
                valFreqMap[val*2] -= cnt
                if valFreqMap[val*2] == 0 {
                    delete(valFreqMap, val*2)
                }
            }
        }
    }
    return res
}