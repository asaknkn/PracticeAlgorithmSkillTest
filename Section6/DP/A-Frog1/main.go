package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var N, ans int

func main() {
	sc.Split(bufio.ScanWords)
	N = nextInt()

	h := make([]int, N)
	for i := 0; i < N; i++ {
		h[i] = nextInt()
	}

	cost := make([]int, N)
	cost[1] = cost[0] + int(math.Abs(float64(h[0]-h[1])))
	for i := 2; i < N; i++ {
		cost[i] = int(math.Min(float64(cost[i-1]+int(math.Abs(float64(h[i-1]-h[i])))), float64(cost[i-2]+int(math.Abs(float64(h[i-2]-h[i]))))))
	}

	fmt.Println(cost[N-1])
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
