/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	var letters []rune 
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}

		if unicode.IsDigit(r) {
			letters = append(letters, r)
		}
	}

	for i := range letters {
		if letters[i] != letters[len(letters) - i - 1] {
			return false
		}
	}

	return true
}
// @lc code=end

