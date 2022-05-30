
// augmented binary search tree approach
// time: O(logn)
// space: O(n)


func min(a, b int) int {if a <= b {return a}; return b}
func max(a, b int) int {if a >= b {return a}; return b}

func overlaps(node *llrbNode, start, end int) bool {
    return max(node.start, start) < min(node.end, end)
}


type MyCalendar struct {
    root *llrbNode
}


func Constructor() MyCalendar {
    return MyCalendar{}
}


func (this *MyCalendar) Book(start int, end int) bool {
    if this.root == nil {
        this.root = &llrbNode{start:start, end: end}
        return true
    }
    var canBook bool
    this.root, canBook = insertHelper(this.root, start, end)
    if !canBook {
        return false
    }
    return true
}




type llrbNode struct {
    start, end int
    left, right *llrbNode
    isRed bool
}
func isRedFn(node *llrbNode) bool {
    if node == nil {
        return false
    }
    return node.isRed
}

func flipColors(x *llrbNode) {
    x.isRed = true 
    x.left.isRed = false
    x.right.isRed = false
}

func rotateRight(x *llrbNode) *llrbNode {
    y := x.left
    x.left = y.right
    y.right = x
    y.isRed = y.right.isRed
    y.right.isRed = true
    return y
}

func rotateLeft(x *llrbNode) *llrbNode {
    y := x.right
    x.right = y.left
    y.left = x
    y.isRed = y.left.isRed
    y.left.isRed = true 
    return y
}

func insertHelper(node *llrbNode, start int, end int) (*llrbNode, bool) {
    if node == nil {
        newChildNode := &llrbNode{start: start, end: end, isRed: true}
        return newChildNode, true
    }
    if overlaps(node, start, end) {
        return node, false
    }
    var canInsert bool
    if node.start < start {
        node.right, canInsert = insertHelper(node.right, start, end)
    } else {
        node.left, canInsert = insertHelper(node.left, start, end)
    }
    // rebalance nodes
    // tree is right heavy
    if isRedFn(node.right) && !isRedFn(node.right.right) {
        node = rotateLeft(node)
    }
    // tree is left heavy
    if isRedFn(node.left) && isRedFn(node.left.left) {
        node = rotateRight(node)
    }
    if isRedFn(node.left) && isRedFn(node.right) {
        flipColors(node)
    }
    return node, canInsert 
}







/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */