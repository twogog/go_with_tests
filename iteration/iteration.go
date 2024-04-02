package iteration

func Repeat(str string, times int) string {
	var result string
	for i := 0; i < times; i++ {
		result += str
	}
	return result
}
