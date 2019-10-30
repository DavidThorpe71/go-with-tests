package iteration

// Repeat returns a given string repeated a given number of times
func Repeat(character string, count int) (repeated string) {
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
