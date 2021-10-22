//time: O(len(l1)+len(l2)) ~ O(n)
//space: O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    carry := 0
    currNode := dummy
    for l1 != nil || l2 != nil {
        val := carry
        if l1 != nil {
            val += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            val += l2.Val
            l2 = l2.Next
        }
        carry, val = val / 10, val % 10
        currNode.Next = &ListNode{Val:val}
        currNode = currNode.Next
    }
    if carry == 1 { currNode.Next = &ListNode{Val:carry}}
    return dummy.Next
    
}