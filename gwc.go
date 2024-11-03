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
		counts := process(os.Stdin, counters...)
		fmt.Println(formatCounts(counts))
		return
	}

	total := make([]int, len(counters))

	for _, fileName := range files {
		file, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
		if err != nil {
			fmt.Println("file " + fileName + " is not valid")
			return
		}

		counts := process(file, counters...)
		for i := range total {
			total[i] += counts[i]
		}
		fmt.Println(formatCounts(counts), fileName)
	}

	if len(files) > 1 {
		fmt.Println(formatCounts(total), "total")
	}
}

func process(r io.Reader, counters ...Counter) []int {
	text, err := io.ReadAll(r)
	if err != nil {
		fmt.Println(err)
	}

	result := make([]int, len(counters))
	for i, c := range counters {
		result[i] = c.Count(text)
	}

	return result
}

func formatCounts(counts []int) string {
	output := ""
	for _, count := range counts {
		output += fmt.Sprintf("%8d", count)
	}
	return output
}

func main() {
	charMode := ""
	isLine := flag.Bool("l", false, "get number of lines")
	isWord := flag.Bool("w", false, "get number of words")
	flag.BoolFunc("c", "get number of bytes", func(string) error { charMode = "c"; return nil })
	flag.BoolFunc("m", "get number of characters", func(string) error { charMode = "m"; return nil })

	flag.Parse()
	counters := make([]Counter, 0)

	if *isLine {
		counters = append(counters, counter.Line())
	}

	if *isWord {
		counters = append(counters, counter.Word())
	}

	switch charMode {
	case "c":
		counters = append(counters, counter.Byte())
	case "m":
		counters = append(counters, counter.Character())
	}

	if len(counters) == 0 {
		processAll(flag.Args(), counter.Line(), counter.Word(), counter.Byte())
	} else {
		processAll(flag.Args(), counters...)
	}
}
