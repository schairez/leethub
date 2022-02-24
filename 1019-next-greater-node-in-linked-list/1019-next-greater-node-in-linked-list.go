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
        monoDecrStack = make([]pair, 0)
        // constraint: 1 <= n <= 10^4
        res           = make([]int, 10000)
        idx           = -1
    )
    var top pair
    for head != nil {
        idx++
        currVal := head.Val
        for len(monoDecrStack)!= 0 && currVal > monoDecrStack[len(monoDecrStack)-1].val {
            lenStack := len(monoDecrStack)
            top, monoDecrStack = monoDecrStack[lenStack-1], monoDecrStack[:lenStack-1]
            res[top.idx] = currVal
        } 
        monoDecrStack = append(monoDecrStack, pair{currVal, idx})
        head = head.Next
    }
    for len(monoDecrStack)!= 0 {
        lenStack := len(monoDecrStack)
        top, monoDecrStack = monoDecrStack[lenStack-1], monoDecrStack[:lenStack-1]
        res[top.idx] = 0
    }
    return res[:idx+1]
}

