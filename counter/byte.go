package counter

type byteCounter struct {
}

func Byte() *byteCounter {
	return &byteCounter{}
}

func (b *byteCounter) Count(text []byte) int {
	return len(text)
}
