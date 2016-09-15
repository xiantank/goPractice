package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	const (
		FLAG_CASEINSENSITIVE = 1 << iota
		FLAG_LINENUM         = 1 << iota
		FLAG_COUNTSUM        = 1 << iota
		FLAG_INVERSE         = 1 << iota
	)
	args := os.Args[1:]
	flag := 0
	var pattern string
	var filename string
	argCur := 0
	if args[argCur][0] == '-' {
		flagString := args[argCur][1:]
		for i := 0; i < len(flagString); i++ {
			switch flagString[i] {
			case 'i':
				flag |= FLAG_CASEINSENSITIVE
			case 'n':
				flag |= FLAG_LINENUM
			case 'c':
				flag |= FLAG_COUNTSUM
			case 'v':
				flag |= FLAG_INVERSE
			}
		}
		argCur++
	}
	pattern = args[argCur]
	argCur++
	filename = args[argCur]
	argCur++
	fmt.Println(pattern, filename)

	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	myReader := bufio.NewReader(fi)
	scanner := bufio.NewScanner(myReader)
	if (flag & FLAG_CASEINSENSITIVE) != 0 {
		pattern = "(?i)" + pattern

	}
	patternRegex := regexp.MustCompile(pattern)

	lineNum := 0
	matchCount := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		matched := patternRegex.MatchString(line)
		if (flag&FLAG_INVERSE == 0) == matched {
			matchCount++
			if flag&FLAG_COUNTSUM == 0 {
				if (flag & FLAG_LINENUM) != 0 {
					fmt.Print(lineNum)
				}
				fmt.Println(line)
			}
		}
	}
	if flag&FLAG_COUNTSUM != 0 {
		fmt.Println(matchCount)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error", err)
	}
}
