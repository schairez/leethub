import "math/rand"

//time: O(1) avg; O(n) when resizeable array resizes
//space: O(n)

type RandomizedSet struct {
    //maps val to idx
    numToIdx map[int]int
    //maps idx to val
    elems []int
}


func Constructor() RandomizedSet {
    return RandomizedSet{ make(map[int]int), []int{} }
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, exists := this.numToIdx[val]; exists {
        return false
    }
    n := len(this.elems)
    this.elems = append(this.elems, val)
    this.numToIdx[val] = n
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    if _, exists := this.numToIdx[val]; !exists {
        return false
    }
    lastIdx := len(this.elems) - 1
    lastElem := this.elems[lastIdx]
    idxLoc := this.numToIdx[val]
    this.elems[idxLoc] = lastElem
    this.numToIdx[lastElem] = idxLoc
    this.elems = this.elems[:lastIdx]
    delete(this.numToIdx, val)
    return true
}


func (this *RandomizedSet) GetRandom() int {
    idx := rand.Intn(len(this.elems))
    return this.elems[idx]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */