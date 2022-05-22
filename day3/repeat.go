package iteration

const repeatCount = 5

// Repeat returns character repeated 5 times.
func Repeat(value string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + value
	}
	return repeated
}
