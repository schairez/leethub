
// 2126. Destroying Asteroids
// time: O(nlogn) to sort
// space: O(logn) to sort

func asteroidsDestroyed(mass int, asteroids []int) bool {
    sort.Ints(asteroids) //time:O(nlogn)
    currMass := mass
    for i := range asteroids {
        planetMass := asteroids[i]
        if currMass < planetMass {
            return false
        }
        currMass += planetMass
    }
    return true
}