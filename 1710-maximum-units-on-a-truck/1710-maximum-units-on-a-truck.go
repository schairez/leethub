//time: O(nlongn + n) ~ O(nlogn)
//space: O(logn) to sort

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
    sort.Slice(boxTypes, func(i, j int) bool { //time: O(nlogn) 
        return boxTypes[i][1] > boxTypes[j][1]
    })
    boxesLeft := truckSize
    totalUnits := 0
    for i := range boxTypes { //time: O(n)
        if boxesLeft == 0 { break }
        numB, numU := boxTypes[i][0], boxTypes[i][1]
        numB = min(numB, boxesLeft)
        totalUnits += (numB*numU)
        boxesLeft -= numB
    }
    return totalUnits
    
}

func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}