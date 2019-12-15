package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	inputs, err := getInput("../input")
	if err != nil {
		log.Fatal(err)
	}

	for _, data := range inputs {
		println(data)
	}
}

func getInput(path string) (lines []string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	return
}
