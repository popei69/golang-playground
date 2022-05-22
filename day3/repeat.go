package iteration

// Repeat returns character for a given time.
func Repeat(value string, counter int) string {
	if counter <= 0 {
		return ""
	}

	var repeated string
	for i := 0; i < counter; i++ {
		repeated += value
	}
	return repeated
}
