package stringTag

/**
520. 检测大写字母
 */

func detectCapitalUse(word string) bool {
	if word == "" {
		return false
	}
	l := len(word)
	if word[0] >= 'A' && word[0] <= 'Z' {
		if l >= 2 {
			if word[1] >= 'A' && word[1] <= 'Z' {
				for i := 2; i < l; i++ {
					if word[i] < 'A' || word[i] > 'Z' {
						return false
					}
				}
			} else {
				for i := 2; i < l; i++ {
					if word[i] >= 'A' && word[i] <= 'Z' {
						return false
					}
				}
			}
		}
	} else {
		for i := 1; i < l; i++ {
			if word[i] >= 'A' && word[i] <= 'Z' {
				return false
			}
		}
	}
	return true
}
