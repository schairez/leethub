import "sort"

//time: O(nlogn + n) ~ O(nlogn)
//space: O(n)
/*
testCases
[6,3,0,1]
[0,0,0,0]

*/


func findOriginalArray(changed []int) []int {
    n := len(changed)
   	if n & 1 == 1 {
		return []int{}
	}
	freqMap := map[int]int{}
	for _, num := range changed {
		freqMap[num]++
	}
    //for testcases with duplicates
   //e.g. [0,0,0,0]
    nums := make([]int, 0, len(freqMap))
    for k := range freqMap {
        nums = append(nums, k)
    }
	sort.Ints(nums)
	ans := make([]int, 0)
	for _, num := range nums {
		if freqMap[2*num] < freqMap[num] {
			return []int{}
		}
		for i := 0; i < freqMap[num]; {
			ans = append(ans, num)
			i++
			freqMap[2*num]--
		}
	}
    
	return ans 
    
}