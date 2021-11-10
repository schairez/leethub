import "math"

var maxSize = int(math.Pow(10, 6) + 1) 


type MyHashSet struct {
    tbl []bool
}


func Constructor() MyHashSet {
    return MyHashSet{make([]bool, maxSize)}
}


func (this *MyHashSet) Add(key int)  {
    this.tbl[key] = true
}


func (this *MyHashSet) Remove(key int)  {
    this.tbl[key] = false
}


func (this *MyHashSet) Contains(key int) bool {
    return this.tbl[key]
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */