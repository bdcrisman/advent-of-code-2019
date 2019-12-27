package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"sync"
)

// SumMutex is a mutex wrapper for the overall sum.
type SumMutex struct {
	sum int
	mux sync.Mutex
}

// AddToSum safely adds to the overall sum.
func (s *SumMutex) AddToSum(value int) {
	s.mux.Lock()
	s.sum += value
	s.mux.Unlock()
}

// I'm utilizing a waitgroup to sum the
// fuel requirements from the input file.
func main() {
	inputArr, err := getInputData("../../input")
	if err != nil {
		log.Fatal(err)
	}

	s := SumMutex{}
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

func getFuelRequirement(mass int, s *SumMutex, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := getAllFuelRequirements(mass, true)
	s.AddToSum(sum)
}

func getAllFuelRequirements(mass int, first bool) int {
	if mass < 0 {
		return 0
	}

	m := int(math.Floor(float64(mass)/3) - 2)
	if first {
		return getAllFuelRequirements(m, false)
	}
	return mass + getAllFuelRequirements(m, false)
}
