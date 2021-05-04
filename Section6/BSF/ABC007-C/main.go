package main

import (
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type point struct {
	x int
	y int
}

func main() {
	sc.Split(bufio.ScanWords)
	r := nextInt()
	c := nextInt()
	sy := nextInt() - 1
	sx := nextInt() - 1
	gy := nextInt() - 1
	gx := nextInt() - 1

	s := make([][]rune, r)
	for i := 0; i < r; i++ {
		ss := next()
		for _, v := range ss {
			s[i] = append(s[i], v)
		}
	}

	dist := make([][]int, r)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			dist[i] = append(dist[i], -1)
		}
	}

	var q []point
	q = append(q, point{x: sy, y: sx})
	dist[sy][sx] = 0

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
