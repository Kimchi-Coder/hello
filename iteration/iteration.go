package iteration

func Repeat(char string, amnt int) (repeated string) {
	for i := 0; i < amnt; i++ {
		repeated += char
	}
	return
}
