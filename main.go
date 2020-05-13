package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var chResult = make(chan largeNumber)
var chInput = make(chan bool)
var input largeNumber
var caches = make(map[string]cache)

type cache struct {
	totalSum largeNumber
}

// Max length of max int
type largeNumber struct {
	numbers []byte
}

var largeNumberOne largeNumber
var largeNumberZero largeNumber

var start time.Time

// Make a 3 GB 'ballast'
// var _ = make([]byte, 10<<30)

func init() {
	largeNumberOne = newLargeNumber("1")
	largeNumberZero = newLargeNumber("0")
}

func main() {
	var answer largeNumber
	// Start concurrent calculation
	go backgroundCalcPowerSum()

	// Get user input concurrently
	go getUserInput()

	// Wait for the result
	answer = <-chResult
	elapsed := time.Since(start)
	fmt.Printf("Time used: %f seconds\n", elapsed.Seconds())

	// Print the result
	// fmt.Printf("Result is: %s\n", answer.String())
	fmt.Println("Result is:")
	answer.Print()
}

func getUserInput() {
	r := bufio.NewReader(os.Stdin)

	// Retrieve input
	fmt.Printf("Enter a non negative number for power sum: ")
	userInputStr, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// Set user input
	setUserInput(newLargeNumber(strings.TrimSpace(userInputStr)))
}

func setUserInput(userInput largeNumber) {
	if userInput.IsLessThan(largeNumberZero) {
		panic("Negative numbers are not allowed:" + userInput.String())
	}

	// atomic.StoreUint64(&input, uint64(userInput))
	input = userInput

	// Signal input received and store
	chInput <- true

	// Start timer after receiving and verifying user input ;D
	start = time.Now()
}

func pow(n, p largeNumber) (a largeNumber) {
	if n.Length() == 1 && n.GetAsInt(0) == 1 {
		return newLargeNumber("1")
	}

	if p.Length() == 1 && p.GetAsInt(0) == 1 {
		return newLargeNumber("1")
	}

	if n.Length() == 1 && n.GetAsInt(0) == 0 {
		return newLargeNumber("0")
	}

	a = newLargeNumber(n.String())
	i := newLargeNumber("1")

	for ; i.IsLessThan(p); i.IncrementByOne() {
		a.Multiply(n)
	}

	return
}

func backgroundCalcPowerSum() {
	// Start of sum is 0
	sum := newLargeNumber("0")
	cur := newLargeNumber("1")
	var hasUserInput bool

	for {
		sum.Add(pow(cur, cur))

		if !hasUserInput {
			caches[cur.String()] = cache{
				totalSum: newLargeNumber(sum.String()),
			}
		}

		cur.IncrementByOne()

		if hasUserInput {
			// Check if user solution is already calculated
			if cur.IsGreaterThan(input) {
				chResult <- caches[input.String()].totalSum
				return
			} else if cur.IsEqual(input) {
				chResult <- sum
			}
		} else {
			select {
			case <-chInput:
				hasUserInput = true
			default:
				// Make sure the goroutine to get user input is running
				runtime.Gosched()
			}
		}
	}
}

func newLargeNumber(s string) largeNumber {
	maxLen := len(s)
	newBytes := make([]byte, maxLen)

	for i, j := 0, maxLen-1; i < maxLen; i++ {
		newBytes[i] = s[j]
		j--
	}

	return largeNumber{numbers: newBytes}
}

func (n *largeNumber) String() string {
	var strBuilder strings.Builder
	for i := n.Length() - 1; i >= 0; i-- {
		strBuilder.WriteByte(n.numbers[i])
	}

	return strBuilder.String()
}

func (n *largeNumber) Append(s string) {
	maxLen := len(s)
	newBytes := make([]byte, maxLen)

	for i, j := 0, maxLen-1; i < maxLen; i++ {
		newBytes[i] = s[j]
		j--
	}

	n.numbers = append(n.numbers, newBytes...)
}

func (n *largeNumber) Replace(s string) {
	maxLen := len(s)
	newBytes := make([]byte, maxLen)

	for i, j := 0, maxLen-1; i < maxLen; i++ {
		newBytes[i] = s[j]
		j--
	}

	n.numbers = newBytes
}

func (n *largeNumber) IsLessThan(c largeNumber) bool {
	// Return true if n is less than c
	if n.Length() < c.Length() {
		return true
	} else if n.Length() > c.Length() {
		return false
	}

	for k := n.Length() - 1; k >= 0; k-- {
		if n.GetAsInt(k) < c.GetAsInt(k) {
			return true
		} else if n.GetAsInt(k) > c.GetAsInt(k) {
			return false
		}
	}

	return false
}

func (n *largeNumber) IsGreaterThan(c largeNumber) bool {
	// Return true if n is greater than c
	if n.IsLessThan(c) || n.IsEqual(c) {
		return false
	}

	return true
}

func (n *largeNumber) IsEqual(c largeNumber) bool {
	if n.IsLessThan(c) || c.IsLessThan(*n) {
		return false
	}

	return true
}

func (n *largeNumber) Get(i int) byte {
	return n.numbers[i]
}

func (n *largeNumber) GetAsInt(i int) int {
	c, _ := strconv.Atoi(string(n.numbers[i]))
	return c
}

func (n *largeNumber) Set(i int, v int) {
	if v >= 10 {
		panic("Input more than 10")
	}

	n.numbers[i] = []byte(strconv.Itoa(v))[0]
}

func (n *largeNumber) Length() int {
	return len(n.numbers)
}

func (n *largeNumber) IncrementByOne() {
	n.Add(largeNumberOne)
}

func (n *largeNumber) Add(l largeNumber) {
	if l.Length() > n.Length() {
		// Swap if input is bigger than instance
		temp := l.String()
		l.Add(*n)
		n.Replace(l.String())
		l.Replace(temp)
		return
	}

	// Add l to n
	// bring forward holder // 0 to 9
	bf := 0
	for k, k1 := 0, 0; k < n.Length(); k++ {
		cur := n.GetAsInt(k)
		var curNew int

		if k1 < l.Length() {
			curNew = l.GetAsInt(k1)
		} else {
			// Add the bring forward if any
			curNew = bf
			bf = 0

			if curNew == 0 {
				break
			}
		}

		if cur+curNew+bf >= 10 {
			newSum := cur + curNew + bf
			newNum := newSum % 10

			n.Set(k, newNum)
			bf = newSum / 10
		} else {
			n.Set(k, cur+curNew+bf)
			bf = 0
		}

		k1++
	}

	if bf > 0 {
		n.Append(strconv.Itoa(bf))
	}
}

func (n *largeNumber) Multiply(l largeNumber) {
	if l.Length() == 0 {
		panic("No input specified in multiplication")
	}

	// Multiply l with n
	// bring forward holder // 0 to 9
	bf := 0
	var newMulti []largeNumber

	for k1 := 0; k1 < l.Length(); k1++ {
		newLN := newLargeNumber("")

		for i := 0; i < k1; i++ {
			newLN.Append("0")
		}

		for k := 0; k < n.Length(); k++ {
			newSum := l.GetAsInt(k1)*n.GetAsInt(k) + bf

			if newSum >= 10 {
				newNum := newSum % 10
				newLN.Append(strconv.Itoa(newNum))
				bf = newSum / 10
			} else {
				newLN.Append(strconv.Itoa(newSum))
				bf = 0
			}
		}
		if bf > 0 {
			newLN.Append(strconv.Itoa(bf))
			bf = 0
		}

		newMulti = append(newMulti, newLN)
	}

	if len(newMulti) >= 1 {
		for k := range newMulti[1:] {
			newMulti[0].Add(newMulti[k+1])
		}

		n.Replace(newMulti[0].String())
	}
}

func (n *largeNumber) Print() {
	var strBuilder strings.Builder
	for i := n.Length() - 1; i >= 0; i-- {
		strBuilder.WriteByte(n.numbers[i])
		if i%3 == 0 {
			strBuilder.WriteRune(',')
		}
	}

	finalStr := strBuilder.String()
	finalStr = finalStr[:len(finalStr)-1]

	fmt.Println(finalStr)
}
