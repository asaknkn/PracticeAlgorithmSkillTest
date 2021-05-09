package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	N := nextInt()
	W := nextInt()

	ws := []int{0}
	vs := []int{0}

	for i := 0; i < N; i++ {
		w := nextInt()
		v := nextInt()
		ws = append(ws, w)
		vs = append(vs, v)
	}

	value := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		for j := 0; j < W+1; j++ {
			value[i] = append(value[i], 0)
		}
	}

	for i := 1; i < N+1; i++ {
		for w := 0; w < W+1; w++ {
			//fmt.Println(value[i])
			value[i][w] = int(math.Max(float64(value[i][w]), float64(value[i-1][w])))

			if w-ws[i] >= 0 {
				value[i][w] = int(math.Max(float64(value[i][w]), float64(value[i-1][w-ws[i]]+vs[i])))
			}
			//fmt.Println(value[i])
		}
		fmt.Println(value[i])
	}

	var ans int
	for w := 0; w < W; w++ {
		ans = int(math.Max(float64(value[N][w]), float64(value[N][w+1])))
		//fmt.Println(ans)
	}
	fmt.Println(ans)
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
