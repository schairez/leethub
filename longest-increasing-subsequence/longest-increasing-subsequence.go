//space optimized patience sort method
//we just need to keep track of the heads of each patience sorted bucket
/*
Input: nums = [10,9,2,5,3,7,101,18]
@idx0
10
@idx1
9
@idx3
2    5

@idx6
2    3   7 101 
after last idx
2    3  7  18

*/

//time: O(nlogn)
//space: O(n)

func lengthOfLIS(nums []int) int {
    n := len(nums)
    if n == 1 { return n }
    var piles []int
    for _, num := range nums {
        idx := binSearchInsertIdx(piles, num)
        if idx == len(piles) {
            piles = append(piles, num)
        }
        piles[idx] = num
        
    }
    return len(piles)
}

/*
make new pile of cards if head of deque < nums[i]
use binSearch to speed this up
//use n deques of piles
//use binary search to find leftmost legal pile

Input: nums = [10,9,2,5,3,7,101,18]
Output: 4

b1  b2  b3  b4
10  5   7   101  
9   3       18
2

note: 4 buckets; therefore len -> 4

*/
///////////////////////////////////////////////////

//patience sorting algo 
//reminiscent of solitaire and of the card game patience
//time: O(nlogn)
//space: O(n^2)
func lengthOfLISPatienceSort(nums []int) int {
    n := len(nums)
    if n == 1 { return n }
    piles := NewPiles()
    for _, val := range nums {
        idx := binSearchInsertIdx(piles.Heads, val)
        if idx == piles.Len() {
            piles.AddNewPile()
        }
        piles.AppendToPile(idx, val)
    }
    return piles.Len()
    
}


type Piles struct {
    PileElems [][]int
    Heads []int
}

func NewPiles() *Piles {
    return &Piles{[][]int{}, []int{} }
}

func (p *Piles) AddNewPile() {
    p.PileElems = append(p.PileElems, []int{})
    p.Heads = append(p.Heads, 0)
}
func (p *Piles) AppendToPile(idx, val int) {
    p.PileElems[idx] = append(p.PileElems[idx], val)
    p.Heads[idx] = val
}
func (p *Piles) Len() int {
    return len(p.Heads)
}

//first element value >= t
func binSearchInsertIdx(heads []int, target int ) int {
    if len(heads) == 0 { return 0 }
    
    lo, hi := 0, len(heads)
    for lo < hi {
        m := lo + (hi - lo) >> 1 
        if heads[m] >= target {
            hi = m
        } else {
            lo = m+1
        } 
    }
    return lo

}




////////////////////////////////////
//DP Version

//1 <= nums.length <= 2500
//state transition fn
//dp[i] stores lis up to i
// dp[i] = max{ dp[i], dp[j]+1 | nums[i] > nums[j] for all j < i}

//time: O(n^2)
//space: O(n)

func max(a, b int) int { if a >= b { return a }; return b}
//iterative
func lengthOfLISDP(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    res := 1
    //dp[i] denotes lis so far up to [0,i]
    //j in [0,i)
    dp[0] = 1
    for i:=1; i < n; i++ {
        dp[i] = 1
        for j:=0; j < i; j++ {
            if nums[i] > nums[j] {
                dp[i] = max(dp[i], dp[j] + 1)
            }
        }
        res = max(dp[i], res)
    }
    
    return res
}


