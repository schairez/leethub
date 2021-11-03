
func canVisitAllRooms(rooms [][]int) bool {
    n := len(rooms)
    curr := 0
    stack := []int{curr}
    visited := make(map[int]bool, n)
    for len(stack) > 0 { //iterative dfs
        curr = stack[len(stack)-1]
        visited[curr] = true
        stack = stack[:len(stack)-1]
        if len(rooms[curr]) != 0 {
            for _, v := range rooms[curr] {
                if !visited[v] { stack = append(stack,v) }
            }
        }
    }
    for i:=0; i < n; i++ {
        if !visited[i] { return false}
    }
    return true
}

//rooms = [[1],[2],[3],[]]
/*
       dfs(0, visited)
        visited[room] = true

*/