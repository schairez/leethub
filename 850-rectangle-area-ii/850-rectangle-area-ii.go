func rectangleArea(rectangles [][]int) int {
	simpleRects := make([][4]int, 0, len(rectangles))
	for _, cur := range rectangles {
		addRectangle([4]int{cur[0], cur[1], cur[2], cur[3]}, &simpleRects, 0)
	}
	sq := 0
	for _, rectangle := range simpleRects {
		sq += (rectangle[2] - rectangle[0]) * (rectangle[3] - rectangle[1])
	}
	return sq % 1_000_000_007
}

func addRectangle(cur [4]int, rectangles *[][4]int, start int) {
	if start == len(*rectangles) {
		*rectangles = append(*rectangles, cur)
		return
	}

	cpart := (*rectangles)[start]
	if cpart[0] >= cur[2] || cpart[1] >= cur[3] || cpart[2] <= cur[0] || cpart[3] <= cur[1] {
		addRectangle(cur, rectangles, start+1)
		return
	}

	if cur[0] < cpart[0] && cpart[0] < cur[2] {
		addRectangle([4]int{cur[0], cur[1], cpart[0], cur[3]}, rectangles, start+1)
	}
	if cur[0] < cpart[2] && cpart[2] < cur[2] {
		addRectangle([4]int{cpart[2], cur[1], cur[2], cur[3]}, rectangles, start+1)
	}
	if cur[1] < cpart[1] && cpart[1] < cur[3] {
		addRectangle([4]int{cur[0], cur[1], cur[2], cpart[1]}, rectangles, start+1)
	}
	if cur[1] < cpart[3] && cpart[3] < cur[3] {
		addRectangle([4]int{cur[0], cpart[3], cur[2], cur[3]}, rectangles, start+1)
	}
}

/*

// 1 <= rectangles.length <= 200
const (
    size = 200
)


type SegTree struct {
    Node [4*size]int // width of active intervals contained in this node
    Tag [4*size]int // if tagVal > 0; this means that the interval contained is active
}
func (segTree *SegTree) pull(left, right, idx int) {
    if segTree.Tag[idx] > 0 {
        segTree.Node[idx] = right - left
        return
    }
    if left + 1 == right {
        segTree.Node[idx] = 0
    } else {
        segTree.Node[idx] = segTree.Node[2*idx+1] + segTree.Node[2*idx+2]
    }
}

func (segTree *SegTree) Update(nodeIdx, openClose, startI, endI, ql, qr int ) {
    if ql <= startI && endI <= qr {
        segTree.Tag[nodeIdx] += openClose 
    } else {
        mid := startI + (endI - startI) >> 1
        lcIdx, rcIdx := nodeIdx*2+1, nodeIdx*2+2
        if qr <= mid {
            segTree.Update(lcIdx, openClose, startI, mid, ql, qr)
        } else if mid <= ql {
            segTree.Update(rcIdx, openClose, mid, endI, ql, qr)
        } else {
            segTree.Update(lcIdx, openClose, startI, mid, ql, qr)
            segTree.Update(rcIdx, openClose, mid, endI, ql, qr)
        }
    }
    segTree.pull(startI, endI, nodeIdx)
}



// segTree Node represents sum over ql and qr ?


func rectangleArea(rectangles [][]int) int {
    const mod = 1e9 + 7
    n := len(rectangles)
    type event struct {x1, x2, y, openClose int}
    events := make([]event, 0, n*2)
    for _, rect := range rectangles {
        x1, y1, x2, y2 := rect[0], rect[1], rect[2], rect[3] 
        events = append(events, event{x1, x2, y1, 1})
        events = append(events, event{x1, x2, y2, -1})
    }
    sort.Slice(events, func(i, j int) bool {
        ev1, ev2 := events[i], events[j]
        return ev1.y < ev2.y
    })
    fmt.Println(events)
    area := 0
    prevY := 0
    segTree := &SegTree{}
    for _, event := range events {
        ql, qr, y, openClose := event.x1, event.x2, event.y, event.openClose
        area += segTree.Node[0] * (y - prevY)
        fmt.Println(area)
        area %= mod
        segTree.Update(0, openClose, 0, size-1, ql, qr)
        prevY = y
    }
    return area 
}
*/


//func (segTree *SegTree) Update(nodeIdx, openClose, startI, endI, ql, qr int ) {











