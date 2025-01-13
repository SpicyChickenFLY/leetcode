package main

// You are given a string s. The score of a string is defined as the sum of the absolute difference between the ASCII values of adjacent characters.
//
// Return the score of s.
//
// 2 <= s.length <= 100
// s consists only of lowercase English letters
func scoreOfString(s string) int {
    score := 0
    for idx := range s {
        if idx == 0 {
            continue
        }
        if (s[idx] >= s[idx - 1]) {
            score += int(s[idx]) - int(s[idx - 1])
        } else {
            score += int(s[idx - 1]) - int(s[idx])
        }
    }
    return score
}
