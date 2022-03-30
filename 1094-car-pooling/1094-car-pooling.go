
// 1094. Car Pooling
// time: sort takes O(nlogn) + O(2n) to iter tripData elems ≈ O(nlogn)
// space: O(logn) to sort + O(2n) to cache tripData elems ≈ O(n)
func carPooling(trips [][]int, capacity int) bool {
    n := len(trips)
    type pair struct {numP, time int}
    tripData := make([]pair, 0, 2*n )
    for _, trip := range trips {
        tripData = append(tripData, pair{trip[0], trip[1]})
        tripData = append(tripData, pair{-trip[0], trip[2]})
    }
    sort.Slice(tripData, func(i, j int) bool {
        if tripData[i].time == tripData[j].time {
            return tripData[i].numP < tripData[j].numP
        }
        return tripData[i].time < tripData[j].time
    })
    currPool := 0
    for i := range tripData {
        currPool += tripData[i].numP
        if currPool > capacity {
            return false
        }
    }
    return true
    
}