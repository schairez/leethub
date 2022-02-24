/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func nextLargerNodes(head *ListNode) []int {
    type pair struct {val, idx int}
    var (
        monoStack          = make([]pair, 0)       // mono decr stack
        res                = make([]int, 10000)    // constraint: 1 <= n <= 10^4
        idxSoFar           = -1
    )
    var top pair
    for head != nil {
        idxSoFar++
        currVal := head.Val
        for stackSize := len(monoStack);
        stackSize != 0 && currVal > monoStack[stackSize-1].val;
        stackSize = len(monoStack) {
            top, monoStack = monoStack[stackSize-1], monoStack[:stackSize-1]
            res[top.idx] = currVal
        } 
        monoStack = append(monoStack, pair{currVal, idxSoFar})
        head = head.Next
    }
    res = res[:idxSoFar+1]
    for stackSize := len(monoStack); stackSize != 0; stackSize = len(monoStack) {
        top, monoStack = monoStack[stackSize-1], monoStack[:stackSize-1]
        res[top.idx] = 0
    }
    return res
}

