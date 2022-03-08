// binary search approach
// time: O(nlogn) to sort + O(min(n1,n2)log(max(n1,n2)) to search for vals
// space: logn + O(maxNumValsIntersect)

func intersection(nums1 []int, nums2 []int) []int {
    binSearch := func(nums []int, target int) bool {
        start, end := 0, len(nums)-1
        for start <= end {
            mid := start + (end - start) >> 1
            if nums[mid] == target {
                return true
            } else if nums[mid] < target {
                start = mid+1
            } else {
                end = mid-1
            } 
        }
        return false
    } 
    
    // sort the largest to do binsearch on this sequence
    n1, n2 := len(nums1), len(nums2)
    maxNumValsIntersect := n2
    if n2 > n1 {
        return intersection(nums2, nums1)
    }
    sort.Ints(nums1)
    valsSet := make(map[int]struct{}, maxNumValsIntersect)
    for _, val := range nums2 {
        if isIntersectVal := binSearch(nums1, val); isIntersectVal {
            valsSet[val] = struct{}{}
        }
    }
    res := make([]int, 0, len(valsSet))
    for key := range valsSet {
            res = append(res, key)
    }
    return res
}


// sort and two pointers approach
// time: O(2nlogn) + O(min(n1, n2)) ≈ O(nlogn)
// space: O(logn) for recursive call stack in sorting std logic
func intersectionV2(nums1 []int, nums2 []int) []int {
    n1, n2 := len(nums1), len(nums2)
    maxNumValsIntersect := n1
    if n2 < n1 { //max intersection vals
        maxNumValsIntersect  = n2
    }
    sort.Ints(nums1) //nlogn time
    sort.Ints(nums2) //nlogn time
    start1, start2 := 0, 0 
    res := make([]int, 0, maxNumValsIntersect)
    for start1 < n1 && start2 < n2 {
        if nums1[start1] < nums2[start2] {
            start1++
        } else if nums2[start2] < nums1[start1] {
            start2++
        } else { // nums1[start1] == nums2[start2] {
            val := nums1[start1]
            res = append(res, val)
            start1++
            start2++
            for start1 < n1 && nums1[start1] == val {
                start1++
            }
            for start2 < n2 && nums2[start2] == val {
                start2++
            }
        }
    }
    return res
    
}

// hashset approach
// time: O(n1 + n2) ≈ O(n)
// space: O(max(n1, n2))
func intersectionV3(nums1 []int, nums2 []int) []int {
    hashSet := make(map[int]struct{})
    for idx := range nums1 {
        hashSet[nums1[idx]] = struct{}{}
    }
    var res []int
    for idx := range nums2 {
        if _, exists := hashSet[nums2[idx]]; exists {
            res = append(res, nums2[idx])
            delete(hashSet, nums2[idx])
        }
    }
    return res
}