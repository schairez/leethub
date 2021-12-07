/*
Input: nums = [2,2,3,3,1]
Output: 1

XOR approach
2 ^ 2 -> 0
0 ^ 3 -> 3
3 ^ 3 -> 0
0 ^ 1 -> 1

note works for any order:
nums = [1,2,2,3,3]
1 ^ 2 -> 3
0001
0010
----
0011

3 ^ 2 -> 1 

0011
0010
-----
0001

1 ^ 3 -> 2 

0001
0011
-----
0010

2 ^ 3 -> 1
*/
//time: O(n)
//space: O(1)
func singleNumber(nums []int) int {
    res := nums[0]
    for i:=1; i < len(nums); i++ {
        res ^= nums[i]
    }
    return res
}