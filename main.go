package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
)

var chResult = make(chan uint64)
var chInput = make(chan bool)
var input uint64
var caches = make([]cache, 1)

type cache struct {
	totalSum uint64
}

// Make a 3 GB 'ballast'
// var _ = make([]byte, 10<<30)

func main() {
	var answer uint64
	// Start concurrent calculation
	go backgroundRunner()

	// Get user input concurrently
	go getUserInput()

	// Wait for the result
	answer = <-chResult

	// Print the result
	fmt.Printf("Result is: %d\n", answer)
}

func getUserInput() {
	var userInput int
	r := bufio.NewReader(os.Stdin)

	// Retrieve input
	fmt.Printf("Enter a non negative number for power sum: ")
	userInputStr, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// Set user input
	userInput, err = strconv.Atoi(strings.TrimSpace(userInputStr))
	setUserInput(userInput)
}

func setUserInput(userInput int) {
	if userInput < 0 {
		panic("Negative numbers are not allowed:" + strconv.Itoa(userInput))
	}

	if userInput >= math.MaxInt64 {
		panic("Input number too large.")
	}

	atomic.StoreUint64(&input, uint64(userInput))

	// Signal input received and store
	chInput <- true
}

func pow(n, p uint64) (a uint64) {
	if n == 1 || p == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}

	a = n
	var i uint64 = 1

	for ; i < p; i++ {
		a *= n
	}

	return
}

func backgroundRunner() {
	backgroundCalcPowerSum()
	calcPowerSum()
}

func backgroundCalcPowerSum() {
	// Start of sum is 0
	var sum uint64 = 0
	var cur uint64 = 1

	for {
		sum += pow(cur, cur)

		caches = append(caches, cache{
			totalSum: sum,
		})

		cur++

		select {
		case <-chInput:
			return
		default:
			// Make sure the goroutine to get user input is running
			runtime.Gosched()
		}
	}
}

func calcPowerSum() {
	// Check if user solution is already calculated
	if len(caches) >= int(input) {
		chResult <- caches[input].totalSum
		return
	}

	// Start from last calculated power sum up to user input
	lastPowerSumNum := len(caches) - 1
	sum := caches[lastPowerSumNum].totalSum

	for i := uint64(lastPowerSumNum) + 1; i <= input; i++ {
		sum += pow(i, i)
	}

	chResult <- sum
}
