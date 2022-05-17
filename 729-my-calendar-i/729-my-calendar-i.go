type MyCalendar struct {
    Activity [][]int
}


func Constructor() MyCalendar {
    return MyCalendar{}
}


func (this *MyCalendar) Book(start int, end int) bool {
    for _, v := range this.Activity {
        if v[0] < end && start < v[1] {
            return false
        }
    }
    this.Activity = append(this.Activity, []int{start, end})
    return true
}


/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */