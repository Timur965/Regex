package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFile(filename string, pattern *regexp.Regexp) []string {
	text, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(text), "\n")
	result := make([]string, 0)

	for i := 0; i < len(lines); i++ {
		lines[i] = strings.TrimRight(lines[i], "\r")
		if len(pattern.FindAllStringSubmatch(lines[i], 1)) != 0 {
			lines[i] = lines[i][:len(lines[i])-1]
			result = append(result, lines[i]+parsing(lines[i]))
		}
	}

	return result
}

func parsing(expression string) string {
	pattern := regexp.MustCompile(`([\+\-\*\/])`)
	index := pattern.FindIndex([]byte(expression))

	left, _ := strconv.ParseInt(string(expression[:index[0]]), 10, 32)
	sign := string(expression[index[0]])
	right, _ := strconv.ParseInt(string(expression[index[0]+1:len(expression)-1]), 10, 32)
	result := ""

	switch sign {
	case "+":
		result = strconv.FormatInt(left+right, 10)
	case "-":
		result = strconv.FormatInt(left-right, 10)
	case "*":
		result = strconv.FormatInt(left*right, 10)
	case "/":
		result = strconv.FormatInt(left/right, 10)
	}

	return result
}

func writeFile(filename string, linesForWrite []string) {
	content := make([]byte, 0)

	for i := 0; i < len(linesForWrite); i++ {
		content = append(content, []byte(linesForWrite[i] + "\n")...)
	}

	err := ioutil.WriteFile(filename, content, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	pattern := regexp.MustCompile(`^[\d]+([\+\-\*\/])[\d]+=[?]$`)

	fmt.Print("Введите имя входного файла: ")
	inputFileName, _, err := reader.ReadLine()

	if err != nil {
		log.Fatal(err)
	}

	expressions := readFile(string(inputFileName), pattern)

	fmt.Print("Введите имя выходного файла: ")
	outputFileName, _, err := reader.ReadLine()

	if err != nil {
		log.Fatal(err)
	}

	writeFile(string(outputFileName), expressions)
}
