// 452. Minimum Number of Arrows to Burst Balloons
// time: O(nlogn + n) ≈ O(nlogn)
// space: O(logn) for sort

func findMinArrowShots(points [][]int) int {
    sort.Slice(points, func(a, b int) bool {
        return points[a][1] < points[b][1]
    })
    numArrows := 1
    // last possible arrow x location = xEnd of the 1st balloon 
    arrowLoc := points[0][1]
    for _, point := range points {
        start, end := point[0], point[1]
        if arrowLoc < start {
            arrowLoc = end
            numArrows++
        }
    }
    return numArrows
}