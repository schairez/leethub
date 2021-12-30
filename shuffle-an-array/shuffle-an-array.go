import "math/rand"

//time: O(n) space and time for shuffle fn

type Solution struct {
    orig []int
}


func Constructor(nums []int) Solution {
    return Solution{ nums }
}


func (this *Solution) Reset() []int {
    return this.orig
}


func (this *Solution) Shuffle() []int {
    rand.Seed(time.Now().UnixNano())
    tmp := make([]int, len(this.orig))
    copy(tmp, this.orig)
    rand.Shuffle(len(tmp), func(i, j int) {
        tmp[i], tmp[j] = tmp[j], tmp[i]
    })
    return tmp
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */