package inversecaptcha

//SumOfAllDigitsThatMatchNext sum sequence of consecutive numbers
func InverseCaptcha(sequence string) int {
	if len(sequence) < 2 {
		return 0
	}

	currentDigit := sequence[0]
	sum := 0
	if sequence[0] == sequence[len(sequence)-1] {
		sum += int(sequence[0] - '0')
	}

	for i := 1; i < len(sequence); i++ {

		if currentDigit == sequence[i] {
			sum += int(sequence[i] - '0')
		} else {
			currentDigit = sequence[i]
		}
	}
	return sum
}
