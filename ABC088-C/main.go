package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	c := make([][]int, 3)
	for i := 0; i < 3; i++ {
		c[i] = nextInts(3)
	}

	ok := true

	if c[0][0]-c[0][1] != c[1][0]-c[1][1] || c[1][0]-c[1][1] != c[2][0]-c[2][1] {
		ok = false
	}

	if c[0][1]-c[0][2] != c[1][1]-c[1][2] || c[1][1]-c[1][2] != c[2][1]-c[2][2] {
		ok = false
	}

	if c[0][0]-c[1][0] != c[0][1]-c[1][1] || c[0][1]-c[1][1] != c[0][2]-c[1][2] {
		ok = false
	}

	if c[1][0]-c[2][0] != c[1][1]-c[2][1] || c[1][1]-c[2][1] != c[1][2]-c[2][2] {
		ok = false
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

// -*-*-*-*-*-*-*-*-
// * I/O utilities *
// -*-*-*-*-*-*-*-*-

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	a, _ := strconv.Atoi(next())
	return a
}

func nextFloat64() float64 {
	a, _ := strconv.ParseFloat(next(), 64)
	return a
}

func nextInts(n int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = nextInt()
	}
	return ret
}
func nextFloats(n int) []float64 {
	ret := make([]float64, n)
	for i := 0; i < n; i++ {
		ret[i] = nextFloat64()
	}
	return ret
}
func nextStrings(n int) []string {
	ret := make([]string, n)
	for i := 0; i < n; i++ {
		ret[i] = next()
	}
	return ret
}

func chars(s string) []string {
	ret := make([]string, len([]rune(s)))
	for i, v := range []rune(s) {
		ret[i] = string(v)
	}
	return ret
}
