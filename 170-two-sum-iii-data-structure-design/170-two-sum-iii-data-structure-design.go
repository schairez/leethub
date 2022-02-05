//170. Two Sum III - Data structure design

type TwoSum struct {
    hashTable map[int]int
}

func Constructor() TwoSum {
    return TwoSum{hashTable: make(map[int]int)}
}

//time: O(1)
//space: O(numKeys)
func (this *TwoSum) Add(number int)  {
    this.hashTable[number]++
}


//numKeys in hashTable
//time: O(numKeys)
//space: O(numKeys)
func (this *TwoSum) Find(value int) bool {
    //since duplicate keys
    //if complement == value; we'd need freq cnts >= 2;
    //if we traverse the whole tbl and fail to find two sum
    //return false
    for key := range this.hashTable {
        complement := value - key
        if freq, ok := this.hashTable[complement]; ok {
            if complement == key && freq < 2 {
                continue
            }
            return true
        }
    }
    return false
}


/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */