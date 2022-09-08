package iteration

func Repeat(ch string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += ch
	}
	return repeated
}
