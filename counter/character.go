package counter

type characterCounter struct {
}

func Character() *characterCounter {
	return &characterCounter{}
}

func (c *characterCounter) Count(text []byte) int {
	return len([]rune(string(text)))
}
