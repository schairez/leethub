type MyCalendarThree struct {
	date    [][]int
	start   map[int]int
	largest int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{
		date:    nil,
		start:   make(map[int]int),
		largest: 0,
	}
}

func (mct *MyCalendarThree) Book(start int, end int) int {
	mct.date = append(mct.date, []int{start, end})
	for k := range mct.start {
		if k >= start && k < end {
			mct.start[k]++
			if mct.start[k] > mct.largest {
				mct.largest = mct.start[k]
			}
		}
	}
	if _, exist := mct.start[start]; !exist {
		current := 0
		for _, v := range mct.date {
			if v[0] <= start && v[1] > start {
				current++
			}
		}
		mct.start[start] = current
		if current > mct.largest {
			mct.largest = current
		}
	}
	return mct.largest
}
 
/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */