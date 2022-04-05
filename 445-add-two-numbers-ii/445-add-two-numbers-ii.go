/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 445. Add Two Numbers II
// time: O(max(n1, n2)) ≈ O(n)
// space:  O(max(n1 + n2)) ≈ O(n)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var l1Stack, l2Stack []*ListNode 
    addNodesToStack := func(node *ListNode, stack *[]*ListNode) {
        for node != nil {
            *stack = append(*stack, node)
            node = node.Next
        }
    }
    addNodesToStack(l1, &l1Stack)
    addNodesToStack(l2, &l2Stack)
    var (
        pollNode *ListNode
        nextNode, headNode *ListNode
    )
    var carry int
    for len(l1Stack) != 0 || len(l2Stack) != 0 || carry != 0 {
        var sumV int
        if len(l1Stack) != 0 {
            pollNode, l1Stack = l1Stack[len(l1Stack)-1], l1Stack[:len(l1Stack)-1]
            sumV += pollNode.Val
        }
        if len(l2Stack) != 0 {
            pollNode, l2Stack = l2Stack[len(l2Stack)-1], l2Stack[:len(l2Stack)-1]
            sumV += pollNode.Val
        }
        tmp :=  sumV + carry
        nodeVal := tmp % 10
        headNode = &ListNode{Val: nodeVal}
        headNode.Next = nextNode
        nextNode = headNode
        carry = tmp / 10
    }
    
    return headNode
}