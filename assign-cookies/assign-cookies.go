/*
Input: g = [1,2,3], s = [1,1]
If s[j] >= g[i], we can assign the cookie j to the child i, and the child i will be content

maximize numContentKids

time: O(2(nlogn) for sort both arr's + n for iter) ~ O(nlogn)
space: O(1)
[7,8,9,10]
[7,8,9,10]
*/

import "sort"

func findContentChildren(g []int, s []int) int {
    if len(s) == 0 { //no content kids b/c no cookies :/
        return 0
    }
    numKids := len(g)
    numCookies := len(s)
    numContentKids := 0
    sort.Ints(g)
    sort.Ints(s)
    currKidIdx := 0
    currCookieIdx := 0
    for currKidIdx != numKids && currCookieIdx != numCookies {
        if s[currCookieIdx] >= g[currKidIdx] {
            numContentKids++
            currKidIdx++
        }
        currCookieIdx++
    } 
    return numContentKids
}