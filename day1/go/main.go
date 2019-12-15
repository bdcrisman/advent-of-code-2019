package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"sync"
)

// SafeSum is a mutex wrapper for the overall sum.
type SafeSum struct {
	sum int
	mux sync.Mutex
}

// AddToSum safely adds to the overall sum.
func (s *SafeSum) AddToSum(value int) {
	s.mux.Lock()
	s.sum += value
	s.mux.Unlock()
}

// For fun, I'm utilizing a goroutine and waitgroup to sum the
// fuel requirements from the input file.
// I guess I could do this one line at a time and it would probably
// take less time...but I'll have to time my solution.
func main() {
	inputArr, err := getInputData("../input")
	if err != nil {
		log.Fatal(err)
	}

	s := SafeSum{}
	var wg sync.WaitGroup
	wg.Add(len(inputArr))

	for _, input := range inputArr {
		num, err := strconv.Atoi(input)
		if err != nil {
			continue
		}
		go getFuelRequirement(num, &s, &wg)
	}

	wg.Wait()
	println(s.sum)
}

func getInputData(path string) (lines []string, err error) {
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

func getFuelRequirement(mass int, s *SafeSum, wg *sync.WaitGroup) {
	defer wg.Done()
	s.AddToSum(int(mass/3) - 2)
}
