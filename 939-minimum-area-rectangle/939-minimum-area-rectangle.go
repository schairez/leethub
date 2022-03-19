
// 0ms O(1)

//            
func min(a, b int) int { if a <= b { return a}; return b}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

type Point struct {
    x, y int
}

// 
func calcArea(p1, p2 Point) int {
    width := abs(p1.x - p2.x)
    height := abs(p2.y - p1.y)
    return width * height
}


func minAreaRect(points [][]int) int {
    n := len(points)
    if n < 4 {
        return 0
    }
    pointSet := make(map[[2]int]struct{}, n)
    for _, point := range points {
        x, y := point[0], point[1]
        pointSet[[2]int{x,y}] = struct{}{}
    }
    maxInt32 := (1 << 31) -1 //maxInt32
    minArea := maxInt32
    
    for _, p1 := range points {
        p1X, p1Y := p1[0], p1[1] //(1,3)  ; (3,1)
        for _, p2 := range points { //
            p2X, p2Y := p2[0], p2[1]
            if p1X == p2X || p2Y == p1Y {
                continue
            }
            // check p3 and p4 in map
            p3Cand := [2]int{p1X, p2Y}
            p4Cand := [2]int{p2X, p1Y}
            if _, ok := pointSet[p3Cand]; !ok {
                continue
            }
            if _, ok := pointSet[p4Cand]; !ok {
                continue
            }
            pt1 := Point{p1X, p1Y}
            pt2 := Point{p2X, p2Y}
            candArea := calcArea(pt1, pt2)
            minArea = min(minArea, candArea)
        }
    }
    if minArea == maxInt32 {
        return 0
    }
    return minArea
}