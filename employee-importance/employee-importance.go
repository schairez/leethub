/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) int {
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