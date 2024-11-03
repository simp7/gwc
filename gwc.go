package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/simp7/gwc/counter"
)

type Counter interface {
	Count(text []byte) int
}

func processAll(files []string, counters ...Counter) {
	if len(files) == 0 {
		processReader(os.Stdin, "", counters...)
		return
	}

	for _, fileName := range files {
		file, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
		if err != nil {
			fmt.Println("file " + fileName + " is not valid")
			return
		}
		processReader(file, fileName, counters...)
	}
}

func processReader(r io.Reader, explain string, counters ...Counter) {
	text, err := io.ReadAll(r)
	if err != nil {
		fmt.Println(err)
	}

	result := ""
	for _, c := range counters {
		countResult := fmt.Sprintf("%8d", c.Count(text))
		result += countResult
	}

	fmt.Println(result, explain)
}

func main() {
	isLine := flag.Bool("l", false, "get number of lines")
	isWord := flag.Bool("w", false, "get number of words")
	isByte := flag.Bool("c", false, "get number of bytes")
	isChar := flag.Bool("m", false, "get number of characters")
	flag.Parse()
	counters := make([]Counter, 0)

	if *isLine {
		counters = append(counters, counter.Line())
	}

	if *isWord {
		counters = append(counters, counter.Word())
	}

	if *isChar {
		counters = append(counters, counter.Character())
	} else if *isByte {
		counters = append(counters, counter.Byte())
	}

	if len(counters) == 0 {
		processAll(flag.Args(), counter.Line(), counter.Word(), counter.Byte())
	} else {
		processAll(flag.Args(), counters...)
	}
}
