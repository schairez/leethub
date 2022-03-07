
// easy hashset approach

func intersection(nums1 []int, nums2 []int) []int {
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