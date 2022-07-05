
// 85. Maximal Rectangle
// time: O(m*n)
// space: O(n)

func max(a, b int) int {if a >= b { return a}; return b}

func maximalRectangle(matrix [][]byte) int {
    numX, numY := len(matrix), len(matrix[0])
    maxArea := 0
    hist := make([]int, numY+1)
    for x := 0; x < numX; x++ {
        for y := 0; y < numY; y++ {
            if matrix[x][y] == '0' {
                hist[y] = 0
            } else {
                hist[y]++
            }
        }
        maxArea = max(maxArea, rectArea(hist))
    }
    return maxArea
}

func rectArea(hist []int) int {
    var stack []int
    area := 0
    for idx, curr := range hist {
        for {
            if len(stack) == 0 || hist[stack[len(stack)-1]] <= curr {
                break
            }
            var height int
            height, stack = hist[stack[len(stack)-1]], stack[:len(stack)-1]
            width := idx
            if len(stack) != 0 {
                width = idx - stack[len(stack)-1] - 1 
            }
            localArea := width * height
            area = max(area, localArea)
        }
        stack = append(stack, idx)
    }
    return area
}
