func findJudge(n int, trust [][]int) int {
    
    // 1 <= n <= 1000
    // hashtable -> 
    discarded := make(map[int]struct{}, n)
    table := make(map[int]int, n)
    judge := 1
    for _, pair := range trust {
        personA, personB := pair[0], pair[1]
        discarded[personA] = struct{}{}
        table[personB]++
        if _, ok := discarded[personB]; ok {
            continue
        }
        judge = personB
    }
    if _, ok := discarded[judge]; ok {
        return -1
    }
    
    if table[judge] == n-1 {
        return judge
    }
    return -1
}

//trust a -> b


/*

invariants : everyone trusts judge
judge trusts nobody

Input: n = 3, trust = [[1,3],[2,3]]
Output: 3

judge = 1
1->3; j = 3;
2->3; j= 3;
if an edge is like [3,2] 3->2 is false so -1 


/*

    graph := make(map[int]map[int]struct{}, n)
    for _, pair := range trust {
        personA, personB := pair[0], pair[1]
        if _, ok := graph[personA]; !ok {
            graph[personA] = make(map[int]struct{})
        }
        graph[personA][personB] = struct{}{}
    }
    discarded := make(map[int]struct{}, n)
    cnt := 0
    judge := 1
    if len(graph[judge]) > 0 {
        discarded[judge] = struct{}{}
        for _, edgeKey := range graph[judge] {
            judge = edgeKey
            judge = 
        }
    }
    for _, edge := range graph[judge] {
        disc
    }
    

*/