//time:O(n) to iter arr
//space:O(1) 

/*
arr = [1,4,2,5,3]

subarrays are:
[1] [1,4] [1,4,2] [1,4,2,5] [1,4,2,5,3]
[4] [4,2] [4,2,5] [4,2,5,3]
[2] [2,5] [2,5,3]
[5] [5,3]
[3]

for each i; val at i will show up for all subarrays <= i 
and all subarrays >= i
  _
 [1,4, 2,5,3]
  _ |  | | |
  ___  | | |
  ______ | |
  _______| | 
           |
  _________|

all subarrays <= i ---> (i + 1)
all subarrays >= i ---> (n - i)
*/

func sumOddLengthSubarrays(arr []int) int {
    n := len(arr)
    res := 0 
    for i :=0; i<n; i++ {
        //number subarrays containing val @ i 
        numSubArrs := (i+1)*(n-i)
        numOddSubArrs := (numSubArrs + 1)/2 
        //add the contribution factor for the odd num of vals at i
        res += (numOddSubArrs * arr[i])
        
    }
    return res 
}