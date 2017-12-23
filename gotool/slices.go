package gotool

// CheckRuneInRunesV1 checks whether rune in runes.
func CheckRuneInRunesV1(runes []rune, r rune) bool {
	for _, each := range runes {
		if each == r {
			return true
		}
	}
	return false
}

// CheckRuneInRunesV2 checks whether rune in runes.
func CheckRuneInRunesV2(runes []rune, r rune) bool {
	set := make(map[rune]bool)
	for _, v := range runes {
		set[v] = true
	}
	return set[r]
}

// Reverse reverses numbers in int slice;
func Reverse(numbers []int) {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}

// FindLongerAndShorterArray sorts two arrays by their length.
func FindLongerAndShorterArray(arrayA, arrayB []int) (longerArray, shorterArray []int) {
	if len(arrayA) >= len(arrayB) {
		return arrayA, arrayB
	} else {
		return arrayB, arrayA
	}
}
