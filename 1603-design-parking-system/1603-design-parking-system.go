type ParkingSystem struct {
    currSlots [3]int
    maxSlots [3]int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    maxSlots := [3]int{big, medium, small}
    var currSlots [3]int
    return ParkingSystem{currSlots, maxSlots}
}

func (this *ParkingSystem) AddCar(carType int) bool {
    currSize := this.currSlots[carType-1]
    maxSize := this.maxSlots[carType-1]
    if currSize != maxSize {
        this.currSlots[carType-1]++
        return true
    }
    return false
}



/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */