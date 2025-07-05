package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

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
	var sb strings.Builder

	b := make([]byte, 1)
	for {
		_, err := f.Read(b)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal(err)
		}

		if unicode.IsLetter(rune(b[0])) {
			sb.Write(b)
		} else {
			if sb.String() == word {
				count++
			}
			sb.Reset()
		}
	}

	fmt.Println(count)
}

// -----------------------------------------------------------------------------------------------------------

// func main() {
// 	defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()

// 	// read the args
// 	if len(os.Args) < 2 {
// 		log.Fatal("Not enough args")
// 	}
// 	word := os.Args[1]
// 	file := os.Args[2]

// 	// open file
// 	f, err := os.Open(file)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	// count the number of occurrences
// 	count := 0
// 	var sb strings.Builder

// 	b := make([]byte, 1)
// 	r := bufio.NewReader(f) // default size is 4096 bytes
// 	for {
// 		_, err := r.Read(b)
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}
// 			log.Fatal(err)
// 		}

// 		if unicode.IsLetter(rune(b[0])) {
// 			sb.Write(b)
// 		} else {
// 			if sb.String() == word {
// 				count++
// 			}
// 			sb.Reset()
// 		}
// 	}

// 	fmt.Println(count)
// }

// -----------------------------------------------------------------------------------------------------------

// func main() {
// 	defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()

// 	// read the args
// 	if len(os.Args) < 2 {
// 		log.Fatal("Not enough args")
// 	}
// 	word := os.Args[1]
// 	file := os.Args[2]

// 	// open file
// 	f, err := os.Open(file)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	// count the number of occurrences
// 	count := 0
// 	var sb strings.Builder
// 	sb.Grow(32) // preallocate memory for the string builder

// 	b := make([]byte, 1)
// 	r := bufio.NewReader(f)
// 	for {
// 		_, err := r.Read(b)
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}
// 			log.Fatal(err)
// 		}

// 		if unicode.IsLetter(rune(b[0])) {
// 			sb.Write(b)
// 		} else {
// 			if sb.String() == word {
// 				count++
// 			}
// 			sb.Reset()
// 			sb.Grow(32)
// 		}
// 	}

// 	fmt.Println(count)
// }
