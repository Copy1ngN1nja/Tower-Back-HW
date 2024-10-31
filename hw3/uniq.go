package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func remove(lines []string, numFields int, numChars int, notCaseSensitive bool) []string {
	result := make([]string, 0)
	for i := 0; i < len(lines); i++ {
		cntFields, wasSpace := 0, true
		j := 0
		for ; j < len(lines[i]) && cntFields < numFields; j++ {
			if wasSpace && lines[i][j] != ' ' {
				cntFields += 1
			}
			if lines[i][j] != ' ' {
				wasSpace = false
			} else {
				wasSpace = true
			}
		}
		if numFields > 0 {
			for j < len(lines[i]) && lines[i][j] != ' ' {
				j++
			}
			for j < len(lines[i]) && lines[i][j] == ' ' {
				j++
			}
		}
		result = append(result, lines[i][min(j+numChars, len(lines[i])):])
	}
	if notCaseSensitive {
		for i := 0; i < len(result); i++ {
			result[i] = strings.ToLower(result[i])
		}
	}
	return result
}

func makeAns(lines []string, c bool, d bool, u bool, notCaseSensitive bool, numFields int, numChars int) []string {
	compareLines := remove(lines, numFields, numChars, notCaseSensitive)
	result := make([]string, 0)
	n := len(lines)
	switch {
	case c:
		cnt := 1
		for i := 1; i < n; i++ {
			if compareLines[i] == compareLines[i-1] {
				cnt += 1
			} else {
				result = append(result, strconv.Itoa(cnt)+" "+lines[i-1])
				cnt = 1
			}
		}
		result = append(result, strconv.Itoa(cnt)+" "+lines[n-1])
	case d:
		for i := 2; i < n; i++ {
			if compareLines[i] != compareLines[i-1] && compareLines[i-1] == compareLines[i-2] {
				result = append(result, lines[i-1])
			}
		}
		if n >= 2 && compareLines[n-1] == compareLines[n-2] {
			result = append(result, lines[n-1])
		}
	case u:
		for i := 0; i < n-1; i++ {
			if i == 0 && compareLines[i] != compareLines[i+1] || i != 0 && compareLines[i] != compareLines[i-1] && compareLines[i] != compareLines[i+1] {
				result = append(result, lines[i])
			}
		}
		if n >= 2 && compareLines[n-1] != compareLines[n-2] {
			result = append(result, lines[n-1])
		}
	default:
		if n > 0 {
			result = append(result, lines[0])
		}
		for i := 1; i < n; i++ {
			if compareLines[i] != compareLines[i-1] {
				result = append(result, lines[i])
			}
		}
	}
	return result
}

func main() {
	c := flag.Bool("c", false, "enable flag c")
	d := flag.Bool("d", false, "enable flag d")
	u := flag.Bool("u", false, "enable flag u")
	i := flag.Bool("i", false, "enable flag i")
	numFields := flag.Int("f", 0, "how many fields do you want to skip in each string")
	numChars := flag.Int("s", 0, "how many symbols do you want to skip in each string")
	flag.Parse()
	inout := flag.Args()
	fileIn, fileOut := "", ""
	if len(inout) >= 1 {
		fileIn = inout[0]
	}
	if len(inout) == 2 {
		fileOut = inout[1]
	}
	if *c == *d && *d || *c == *u && *u || *d == *u && *u {
		fmt.Println("Failed to run uniq. You can't enable flags -c, -d, -u at the same time")
		return
	}
	lines := make([]string, 0)
	if fileIn != "" {
		file, err := os.Open(fileIn)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			lines = append(lines, line)
		}
	}
	ans := makeAns(lines, *c, *d, *u, *i, *numFields, *numChars)
	if fileOut == "" {
		for _, val := range ans {
			fmt.Println(val)
		}
	} else {
		fileWrite, err := os.Create(fileOut)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, val := range ans {
			fileWrite.Write([]byte(val))
			fileWrite.Write([]byte("\n"))
		}
	}
}
