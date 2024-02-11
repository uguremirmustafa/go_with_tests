package iteration

func Repeat(count int, char string) (result string) {
	for i := 0; i < count; i++ {
		result += char
	}
	return
}
