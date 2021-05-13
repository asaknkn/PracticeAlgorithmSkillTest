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
	N := nextInt()
	var ps []int
	ps = append(ps, 0)
	var P int
	for i := 0; i < N; i++ {
		v := nextInt()
		ps = append(ps, v)
		P += v
	}

	exist := make([][]bool, N+1)
	for i := 0; i < N+1; i++ {
		for j := 0; j < P+1; j++ {
			exist[i] = append(exist[i], false)
		}
	}

	exist[0][0] = true

	for i := 1; i < N+1; i++ {
		for s := 0; s < P+1; s++ {
			if exist[i-1][s] {
				exist[i][s] = true
			}
			if s >= ps[i] && exist[i-1][s-ps[i]] {
				exist[i][s] = true
			}
		}
	}

	var ans int
	for s := 0; s < P+1; s++ {
		if exist[N][s] {
			ans++
		}
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
