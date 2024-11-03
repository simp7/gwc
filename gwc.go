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
	isByte := flag.Bool("c", false, "get number of bytes")
	isLine := flag.Bool("l", false, "get number of lines")
	isWord := flag.Bool("w", false, "get number of words")
	isChar := flag.Bool("m", false, "get number of characters")
	flag.Parse()

	if *isByte {
		processAll(os.Args[2:], counter.Byte())
	} else if *isLine {
		processAll(os.Args[2:], counter.Line())
	} else if *isWord {
		processAll(os.Args[2:], counter.Word())
	} else if *isChar {
		processAll(os.Args[2:], counter.Character())
	} else {
		processAll(os.Args[1:], counter.Line(), counter.Word(), counter.Byte())
	}
}
