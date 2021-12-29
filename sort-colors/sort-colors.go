//time: O(n)
//space: O(1)

//dutch flag algo approach
//sort in-place
//precondition: int vals 0,1,2 only
func sortColors(nums []int)  {
    swap := func(i, j int) {
        nums[i], nums[j] = nums[j], nums[i]
    }
    red, white, blue := 0, 0, len(nums)-1
    for white <= blue {
        switch nums[white] {
        case 0:
            swap(red, white)
            red++
            white++
        case 1:
            white++
        case 2:
            swap(blue, white)
            blue--
        }
    } 
}

/*
//prev version

func sortColors(nums []int)  {
    swap := func(i, j int) {
        nums[i], nums[j] = nums[j], nums[i]
    }
    n := len(nums)
    midVal := 1 //mid vals should be 1
    //two pointer approach
    i, j := 0, 0
    //topIdx captures vals = 2
    topIdx := n-1
    for j <= topIdx {
        switch {
        case nums[j] < midVal:
            swap(i, j)
            i++
            j++
        case nums[j] > midVal:
            swap(j, topIdx)
            topIdx--
        default: //if nums[j] == midVal -> idx j in right pos
            j++
        }
    }
    
}
version 5/29/2021

func sortColors(nums []int)  {
    low, mid, high := 0,0, len(nums)-1
    for mid <= high {
        switch nums[mid] {
            case 0:
                nums[low], nums[mid] = nums[mid], nums[low]
                low++
                mid++
            case 1:
                mid++
            case 2:
                nums[high], nums[mid] = nums[mid], nums[high]
                high--
        }
    }
    
}


*/