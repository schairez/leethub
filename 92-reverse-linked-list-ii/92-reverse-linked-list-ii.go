/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// time: O(n)
// space: O(1)

func reverseBetween(head *ListNode, left int, right int) *ListNode {
    // condition: 1 <= n <= 500
    // 1 <= left <= right <= n
    
    dummyNode := &ListNode{-1, head}
    var (
        parNode, leftNode *ListNode
        nextNode, currNode, tmpNode *ListNode
    )
    parNode = dummyNode
    leftNode = head 
    for i := 0; i < left -1; i++ {
        parNode = leftNode
        if leftNode.Next != nil {
            leftNode = leftNode.Next
        }
    }
    cnt := right - left + 1
    currNode = leftNode
    for {
        tmpNode = currNode.Next
        currNode.Next = nextNode
        nextNode = currNode
        currNode = tmpNode
        cnt--
        if cnt == 0 {
            parNode.Next.Next = currNode
            parNode.Next = nextNode 
            break
        }
    }
    return dummyNode.Next
}