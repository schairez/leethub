//dividend / divisor = quotient

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    res := make([]float64, len(queries))
    divisionGraph := buildGraph(equations, values)
    for idx, query := range queries {
        dividend, divisor := query[0], query[1]
        _, dividendExistsCond := divisionGraph[dividend]
        _, divisorExistsCond := divisionGraph[dividend][divisor]
        if dividendExistsCond && divisorExistsCond {
            res[idx] = divisionGraph[dividend][divisor]
        } else {
            res[idx] = backtrack(dividend, divisor, set{}, divisionGraph) 
        }
    }
    return res
}

type set map[string]bool
type graph map[string]map[string]float64

func buildGraph(equations [][]string, values []float64) graph {
    g := make(graph)
    for idx, eq := range equations {
        dividend, divisor := eq[0], eq[1]
        quotient := values[idx]
        if _, ok := g[dividend]; !ok {
            g[dividend] = map[string]float64{}
        }
        if _, ok := g[divisor]; !ok {
            g[divisor] = map[string]float64{}
        }
        g[dividend][divisor] = quotient
        g[divisor][dividend] = 1.0 / quotient
        g[divisor][divisor] = 1.0
        g[dividend][dividend] = 1.0
    }
    return g
}

func backtrack(dividend, divisor string, visited set, divisionGraph graph) float64 {
    fmt.Println(dividend, divisor)
    visited[dividend] = true
    if _, ok := divisionGraph[divisor]; !ok { //if node not in graph
        return -1.0
    }
    if dividend == divisor { 
        return 1.0 
    }
    var res float64 = -1.0
    for nei := range divisionGraph[dividend] {
        fmt.Println(nei)
        //if nei == dividend { continue } //ignore the self-loops
        if _, seen := visited[nei]; seen {
            continue
        }
        //cumulative product of quotients; e.g. a/b * b/c = a/c
        currQuotient := divisionGraph[dividend][nei]
        nextQuotient := backtrack(nei, divisor, visited, divisionGraph)
        if nextQuotient != -1.0 {
            return currQuotient*nextQuotient
            res = currQuotient * nextQuotient
        }
    }
    
    return res
}