package iteration

const RepeatTimes = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < RepeatTimes; i++ {
		repeated += character
	}
	return repeated
}
