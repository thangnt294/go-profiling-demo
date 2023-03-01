package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

func main() {

	// read the args
	if len(os.Args) < 2 {
		log.Fatal("Not enough args")
	}
	word := os.Args[1]
	file := os.Args[2]

	// open file
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// count the number of occurrences
	count := 0
	curr := ""

	b := make([]byte, 1)
	for {
		_, err := f.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if unicode.IsLetter(rune(b[0])) {
			curr += string(b)
		} else {
			if curr == word {
				count++
			}
			curr = ""
		}
	}

	fmt.Println(count)
}
