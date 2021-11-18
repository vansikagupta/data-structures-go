// leetcode #5: https://leetcode.com/problems/longest-palindromic-substring/submissions/

package strings

//Brute Force

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	maxLen := 1
	palSubstring := s[0:1]

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			substring := s[i : j+1]

			if isPalindrome(substring) && j-i+1 > maxLen {
				palSubstring = substring
				maxLen = j - i + 1
			}
		}
	}

	//fmt.Println(maxLen)
	//fmt.Println(palSubstring)
	return palSubstring

}
