
func min(a, b int) int {if a <= b {return a}; return b}
func max(a, b int) int {if a >= b {return a}; return b}

func overlaps(node *SegNode, start, end int) bool {
    return max(node.start, start) < min(node.end, end)
}

type SegNode struct {
    start, end int
    left, right *SegNode
}


type MyCalendar struct {
    root *SegNode
}


func Constructor() MyCalendar {
    return MyCalendar{}
}


func (this *MyCalendar) Book(start int, end int) bool {
    if this.root == nil {
        this.root = &SegNode{start:start, end: end}
        return true
    }
    if overlaps(this.root, start, end) {
        return false
    }
    if this.root.start < start {
        return this.helper(this.root, this.root.right, start, end)
    } else {
        return this.helper(this.root, this.root.left, start, end)
    }
}

func (this *MyCalendar) helper(parent, childNode *SegNode, start int, end int) bool {
    if childNode == nil {
        node := &SegNode{start: start, end: end}
        if parent.start < start {
            parent.right = node
        } else {
            parent.left = node
        }
        return true
    }
    if overlaps(childNode, start, end) {
        return false
    }
    if childNode.start < start {
        return this.helper(childNode, childNode.right, start, end)
    } else {
        return this.helper(childNode, childNode.left, start, end)
    }
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */