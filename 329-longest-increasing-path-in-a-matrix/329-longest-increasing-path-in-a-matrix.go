func max(a, b int) int {if a >= b { return a}; return b}

func longestIncreasingPath(matrix [][]int) int {
    numR, numC := len(matrix), len(matrix[0])
    dag := make([][]int, numR*numC)
    inDegree := make([]int, numR*numC) 
    dirs := [5]int{1, 0, -1, 0, 1}
    for x := 0; x < numR; x++ {
        for y := 0; y < numC; y++ {
            nodeVal := matrix[x][y]
            srcId := x * numC + y
            for i := 1; i < 5; i++ {
                dX, dY := dirs[i-1], dirs[i]
                nextX, nextY := x + dX, y + dY
                if nextX < 0 || nextX >= numR || nextY < 0 || nextY >= numC {
                    continue
                }
                if matrix[nextX][nextY] <= nodeVal {
                    continue
                }
                dstId := nextX * numC + nextY
                dag[srcId] = append(dag[srcId], dstId)
                inDegree[dstId]++ 
            }
        }
    }
    type nodeData struct {id, dist int}
    var (
        queue []nodeData
        currNode nodeData
    )
    for nodeId, degree := range inDegree {
        if degree == 0 {
            queue = append(queue, nodeData{nodeId,1})
        }
    }
    res := 0
    for len(queue) != 0 {
        currNode, queue = queue[0], queue[1:]
        res = max(res, currNode.dist)
        if len(dag[currNode.id]) > 0 {
            for _, neiId := range dag[currNode.id] {
                if inDegree[neiId] - 1 == 0 {
                    queue = append(queue, nodeData{neiId, currNode.dist+1})
                }
                inDegree[neiId]--
            }
        }
    }
    return res
}