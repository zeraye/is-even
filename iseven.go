package iseven

func IsEven(number int) bool {
	if number == 0 {
		return true
	} else if number < 0 {
		return IsEven(-number)
	}
	return !IsEven(number - 1)
}
