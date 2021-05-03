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
	a := make([][]int, 3)
	for i := 0; i < 3; i++ {
		a[i] = nextInts(3)
	}
	n := nextInt()
	b := nextInts(n)

	M := make([][]bool, 3)
	for i := 0; i < 3; i++ {
		M[i] = make([]bool, 3)
		for j := 0; j < 3; j++ {
			M[i][j] = false
		}
	}

	for _, v := range b {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if a[i][j] == v {
					M[i][j] = true
				}
			}
		}
	}

	bingo := false

	for i := 0; i < 3; i++ {
		if M[i][0] && M[i][1] && M[i][2] {
			bingo = true
		}
	}

	for i := 0; i < 3; i++ {
		if M[0][i] && M[1][i] && M[2][i] {
			bingo = true
		}
	}

	if (M[0][0] && M[1][1] && M[2][2]) || (M[0][2] && M[1][1] && M[2][0]) {
		bingo = true
	}

	if bingo {
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
