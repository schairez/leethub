
func isCouple(num1, num2 int) bool {
    if num1 % 2 == 0 && num2 % 2 == 0 {
        return false
    }
    if num1 % 2 == 1 && num2 % 2 == 1 {
        return false
    }
    if num1 % 2 == 0 && num1 + 1 != num2 {
        return false
    } else if num2 % 2 == 0 && num2 + 1 != num1 {
        return false
    }
    return true
}

func minSwapsCouples(row []int) int {
    n := len(row)
    if n == 2 {
        return 0
    }
    res := 0
    i, j, k := 0, 1, 2
    for i < n-2 {
        if isCouple(row[i], row[j]) {
            i += 2
            j = i + 1
            k = j + 1
        } else {
            for k < n && !isCouple(row[i], row[k]) {
                k++
            }
            row[j], row[k] = row[k], row[j]
            i += 2
            j = i + 1
            k = j + 1
            res++
        }
    }
    return res
}