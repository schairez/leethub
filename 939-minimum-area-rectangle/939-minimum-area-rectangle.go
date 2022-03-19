
func min(a, b int) int { if a <= b { return a}; return b}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

type point struct {
    x, y int
}

// 
func calcArea(p1, p2 point) int {
    width := abs(p1.x - p2.x)
    height := abs(p2.y - p1.y)
    return width * height
}



func minAreaRect(points [][]int) int {
    const maxInt32 = (1 << 31) - 1
    n := len(points)
    if n < 4 {
        return 0
    }
    var minArea int
    pointKeys := make([]point, 0, n)
    pointSet := make(map[int]map[int]struct{}, n)
    //pointSet := make(map[point]struct{}, n)
    for _, elem := range points {
        x, y := elem[0], elem[1]
        pt := point{x,y}
        if _, xExists := pointSet[x]; !xExists {
            pointSet[x] = make(map[int]struct{})
        }
        pointSet[x][y] = struct{}{}
        pointKeys = append(pointKeys, pt)
    }
    
    minArea = maxInt32
    for _, p1 := range pointKeys {
        for _, p2 := range pointKeys {
            if p1 == p2 || p1.x == p2.x || p1.y == p2.y {
                continue
            }
            //p3Cand := point{p1.x, p2.y}
            //p4Cand := point{p2.x, p1.y}
            // for p3
            if _, exists := pointSet[p1.x][p2.y]; !exists {
                continue
            }
            // for p4
            if _, exists := pointSet[p2.x][p1.y]; !exists {
                continue
            }
            candArea := calcArea(p1, p2)
            minArea = min(minArea, candArea)
        } 
    }
    if minArea == maxInt32 {
        return 0
    }
    return minArea
}

