/*
Intuitive answer: loop through letters in O(len(letters))
Also possible: binary search in O(log(len(letters)))
*/

func nextGreatestLetter(letters []byte, target byte) byte {
    if int(letters[len(letters)-1]) <= int(target) {
        return letters[0]
    }
    left := 0
    right := len(letters)-1
    var mid, diff int
    for left < right {
        mid = left + (right - left)/2
        diff = int(letters[mid]) - int(target)
        if diff > 0 {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return letters[left]
}
