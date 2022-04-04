/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 143. Reorder List
// condition: numNodes in range [1, 5*10^4]
// time: O(n)
// space: O(n)

func reorderList(head *ListNode)  {
    var (
        fromHeadNode *ListNode
        fromTailNode *ListNode
        stack []*ListNode
    )
    fromHeadNode = head
    n := 0
    for fromHeadNode != nil {
        stack = append(stack, fromHeadNode)
        fromHeadNode = fromHeadNode.Next
        n++
    }
    fromHeadNode = head
    for numIter := (n-1) >> 1; numIter != 0; numIter-- {
        fromTailNode, stack = stack[len(stack)-1], stack[:len(stack)-1] 
        tmp := fromHeadNode.Next
        fromHeadNode.Next = fromTailNode
        fromTailNode.Next = tmp
        fromHeadNode = tmp
    }
    fromTailNode, stack = stack[len(stack)-1], stack[:len(stack)-1] 
    fromTailNode.Next = nil
    
}