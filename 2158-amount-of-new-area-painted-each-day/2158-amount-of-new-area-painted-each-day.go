
/*

1 <= paint.length <= 10^5
paint[i].length == 2
0 <= starti < endi <= 5 * 104

4*end+1 ?

*/

const (
    size = 5 * 100_000 + 1
)

// reveals if nodeIdx is covered already
type SegTree struct {
    Node [4*size]int
}

func (segTree *SegTree) Upsert(nodeIdx, sIdx, eIdx, ql, qr int) int {
    if sIdx == eIdx || qr <= sIdx || ql >= eIdx {
        return 0
    }
    /*
    if eIdx - sIdx == 1 {
        if segTree.Node[nodeIdx] == 0 {
            segTree.Node[nodeIdx] = 1
            return 1
        }
        return 0
    }
    */
    if segTree.Node[nodeIdx] == eIdx - sIdx {
        return 0
    }
    if ql == sIdx && qr == eIdx {
        prevPainted := segTree.Node[nodeIdx]
        fullSegPaint := eIdx - sIdx
        if prevPainted == fullSegPaint {
            return 0
        }
        newPainted := fullSegPaint - prevPainted
        segTree.Node[nodeIdx] += newPainted 
        return newPainted 
    }
    mid := sIdx + (eIdx - sIdx) >> 1
    var leftQuery, rightQuery int
    // disj cond
    if qr <= mid {
        leftQuery = segTree.Upsert(nodeIdx*2+1, sIdx, mid, ql, qr)
    } else if ql > mid {
        rightQuery = segTree.Upsert(nodeIdx*2+2, mid, eIdx, ql, qr)
    } else {
        leftQuery = segTree.Upsert(nodeIdx*2+1, sIdx, mid, ql, mid)
        rightQuery = segTree.Upsert(nodeIdx*2+2, mid, eIdx, mid, qr)
    }
    segTree.Node[nodeIdx] += leftQuery + rightQuery
    return leftQuery + rightQuery
}

func amountPainted(paint [][]int) []int {
    n := len(paint)
    segTree := &SegTree{}
    res := make([]int, n)
    for i := range paint {
        start, end := paint[i][0], paint[i][1]
        res[i] = segTree.Upsert(0, 0, size, start, end)
    }
    /*
    fmt.Println(segTree.Node[262143])
    fmt.Println(segTree.Node[262144])
    fmt.Println(segTree.Node[262145])
    fmt.Println(segTree.Node[262146])
    */
    return res
}








