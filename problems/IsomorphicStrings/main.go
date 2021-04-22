package main

func isIsomorphic(s string, t string) bool {
	dict := make(map[byte]byte)
	reverse := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if _, ok := dict[s[i]]; !ok {
			// not in dict
			if _, ok := reverse[t[i]]; ok {
				// conflict
				return false
			} else {
				dict[s[i]] = t[i]
				reverse[t[i]] = s[i]
			}
		} else if dict[s[i]] != t[i] {
			return false
		}
	}
	return true
}

func main() {}
