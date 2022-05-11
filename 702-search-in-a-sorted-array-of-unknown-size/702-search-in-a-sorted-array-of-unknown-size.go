/**
 * // This is the ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 * type ArrayReader struct {
 * }
 *
 * func (this *ArrayReader) get(index int) int {}
 */

func search(reader ArrayReader, target int) int {
    res := -1
    lo, hi := 0, 1 
    for reader.get(hi) < target {
        hi *= 2
    }
    for lo + 1 < hi {
        mid := lo + (hi - lo) >> 1
        if reader.get(mid) <= target {
            lo = mid
        } else {
            hi = mid
        }
    }
    
    if reader.get(lo) == target {
        res = lo
    } else if reader.get(hi) == target {
        res = hi
    }
    
    return res
}