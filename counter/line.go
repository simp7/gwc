package counter

type lineCounter struct {
}

func Line() *lineCounter {
	return &lineCounter{}
}

func (l *lineCounter) Count(text []byte) int {
	count := 0
	for _, char := range text {
		if char == '\n' {
			count++
		}
	}
	return count
}
