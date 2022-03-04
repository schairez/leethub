
// 605. Can Place Flowers
// time: O(n)
// space: O(1)

const (
    empty = iota
    planted
)

func canPlaceFlowers(flowerbed []int, n int) bool {
    if n == 0 {
        return true
    }
    numLeft := n
    lastPlantedLoc := -1
    for idx := 0; idx < len(flowerbed); idx++ {
        if numLeft == 0 {
            break
        }
        if flowerbed[idx] == planted {
            continue
        }
        if idx != 0 && flowerbed[idx-1] == planted {
            continue
        } 
        if lastPlantedLoc != -1 && lastPlantedLoc == idx-1 {
            continue
        } 
        if idx+1 != len(flowerbed) && flowerbed[idx+1] == planted {
            continue
        }
        lastPlantedLoc = idx
        numLeft--
    }
    
    return numLeft == 0
}