package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"sync/atomic"
)

var chResult chan uint64
var input uint64
var inputSet int32

func main() {
	go calcPowerSum()
	// fmt.Println(math.MaxUint64)
	var userInput int
	var r bufio.Reader

	fmt.Println("Enter a non negative number for power sum:")
	userInputStr, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	userInput, err = strconv.Atoi(userInputStr)
	setUserInput(userInput)
}

func setUserInput(userInput int) {
	if userInput <= 0 {
		panic("Negative numbers are not allowed")
	}

	if userInput >= math.MaxInt32 {
		panic("Input number too large.")
	}

	atomic.StoreUint64(&input, uint64(userInput))

	atomic.StoreInt32(&inputSet, 1)
}

func pow(n, p uint64) (a uint64) {
	a = n
	var i uint64 = 1

	for ; i < p; i++ {
		a *= n
	}

	return
}

func calcPowerSum() {
	var sum uint64 = 0
	var cur uint64 = 1

	for {
		sum += uint64(pow(cur, cur))
		if cur == 10 {
			fmt.Println(cur, uint64(pow(cur, cur)))
		}
		cur++

		if atomic.LoadInt32(&inputSet) == 1 && atomic.LoadUint64(&input) <= cur {
			chResult <- sum
		} else {
			fmt.Println("Current:", cur-1, "Sum:", sum)
		}
	}
}
