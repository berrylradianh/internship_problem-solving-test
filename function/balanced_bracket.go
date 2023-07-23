package function

// Buat fungsi untuk menemukan Balanced Bracket dengan kompleksitas paling rendah!

func BalancedBracket(s string) string {
	stack := []int32{}

	brackets := map[int32]int32{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	isOpeningBracket := func(c int32) bool {
		_, isOpening := brackets[c]
		return isOpening
	}

	for _, char := range s {
		if isOpeningBracket(char) {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return "NO"
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if brackets[top] != char {
				return "NO"
			}
		}
	}

	if len(stack) == 0 {
		return "YES"
	}

	return "NO"
}
