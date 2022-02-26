
/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

//BFS approach
func getImportance(employees []*Employee, id int) int {
    employeeMap := make(map[int]*Employee, len(employees))
    for _, emp := range employees {
        employeeMap[emp.Id] = emp
    }
    visited := make(map[int]struct{})
    visited[id] = struct{}{}
    var (
        queue   []*Employee
        node    *Employee
        total  int
    )
    
    queue = append(queue, employeeMap[id])
    for len(queue) != 0 {
        node, queue = queue[0], queue[1:]
        total += node.Importance
        visited[node.Id] = struct{}{}
        for _, nei := range node.Subordinates {
            if _, seen := visited[nei]; !seen {
                visited[nei] = struct{}{}
                queue = append(queue, employeeMap[nei])
            }
        }
    }
    return total
}























//time: O(N)
//space: O(N)
func getImportanceDFS(employees []*Employee, id int) int {
    n := len(employees)
    employeeMap := make(map[int]*Employee, n)
    for i := range employees {
        employee := employees[i]
        empId := employee.Id
        employeeMap[empId] = employee
    }
    visitedMap := make(map[int]bool, n)
    var dfs func(empId int) int
    dfs = func(empId int) int {
        employee, ok := employeeMap[empId] 
        if !ok { return 0 }
        visitedMap[empId] = true
        
        path := employee.Importance 
        for _, subId := range employee.Subordinates {
            if !visitedMap[subId] {
                path += dfs(subId)
            }
        }
        return path
    }
    return dfs(id)
}