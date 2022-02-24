/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


// mono decr stack version
// time: O(n)
// space: O(n)
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
    for stackSize := len(monoStack); stackSize != 0; stackSize = len(monoStack) {
        top, monoStack = monoStack[stackSize-1], monoStack[:stackSize-1]
        res[top.idx] = 0
    }
    return res[:idxSoFar+1]
}



// time: O(n^2)
// space: O(1) if not including res output 

func nextLargerNodesV2(head *ListNode)  []int {
    nextGreaterElem := func(node *ListNode, prevVal int) int {
        if node == nil {
            return 0
        }
        for node != nil && prevVal >= node.Val {
            node = node.Next
        }
        if node == nil {
            return 0
        }
        return node.Val
    }
    
    res := make([]int, 10000)    // constraint: 1 <= n <= 10^4
    idxSoFar := -1
    for head != nil {
        idxSoFar++
        res[idxSoFar] = nextGreaterElem(head.Next, head.Val)
        head = head.Next
    }
    return res[:idxSoFar+1]
}

