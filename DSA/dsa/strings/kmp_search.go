package strings

/*
text - abxabcabcaby
pattern - abcaby

a b c a b y
0 0 0 1 2 0 
*/

/*
example pattern

a a b a a b a a a
0 1 0 1 2 3 4 5
*/
func getlps(pattern string) []int {

	lps := make([]int, len(pattern))
	i, j := 0, 1

	for j < len(pattern) {

		if pattern[j] == pattern[i] {
			i++
			lps[j] = i
			j++
		} else {
			if i == 0 {
				j++
			} else {
				i = lps[i-1]
			}
		}
	}
	return lps

}

func KmpSearch(text, pattern string) []int {
	lps := getlps(pattern)
	ans := []int{}
	i, j := 0, 0

	for i < len(text) {

		if text[i] == pattern[j] {
			i++
			j++
		} else if text[i] != pattern[j] {
			if j == 0 {
				i++
			} else {
				j = lps[j-1]
			}
		}

		if j == len(pattern) {
			ans = append(ans, i-len(pattern))
			j = lps[j-1]
		}

	}
	return ans

}
