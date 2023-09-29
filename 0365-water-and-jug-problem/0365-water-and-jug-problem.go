
// n = j1 + j2
// time: O(n)
// space: O(n)

func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
  return bfs(jug1Capacity, jug2Capacity, targetCapacity)
}


func bfs(j1, j2, target int) bool {
  total := j1 + j2
  visited := make([]bool, total+1)
  ops := [4]int{j1, -j1, j2, -j2}
  var (
    queue []int
    node int
  )
  queue = append(queue, 0)
  visited[0] = true
  for len(queue) != 0 {
    node, queue = queue[0], queue[1:]
    if node == target {
      return true
    }
    for _, op := range ops {
      newNode := node + op
      if newNode < 0 || newNode > total {
        continue
      }
      if visited[newNode] {
        continue
      }
      queue = append(queue, newNode)
      visited[newNode] = true
    }
  }
  return false
}






/*

0
vis[0] = t
0 + 3 -> 3
0 + 5 -> 5

q = [3,5]
  2
2 
 


*/