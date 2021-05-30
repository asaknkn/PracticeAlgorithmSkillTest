package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func maxInt(a []int) int {
	sort.Sort(sort.IntSlice(a))
	return a[len(a)-1]
}

func minInt(a []int) int {
	sort.Sort(sort.IntSlice(a))
	return a[0]
}

func rec(i int, lenght []int, done []bool, edges [][]int) int {
	if done[i] {
		return lenght[i]
	}
	lenght[i] = 0
	for _, j := range edges[i] {
		lenght[i] = int(math.Max(float64(lenght[i]), float64(rec(j, lenght, done, edges)+1)))
	}
	done[i] = true
	return lenght[i]
}

func main() {
	sc.Split(bufio.ScanWords)
	N := nextInt()
	M := nextInt()

	edges := make([][]int, N)
	indeg := make([]int, N)

	for i := 0; i < M; i++ {
		x := nextInt()
		y := nextInt()

		edges[x-1] = append(edges[x-1], y-1)
		indeg[y-1] += 1
	}

	length := make([]int, N)
	done := make([]bool, N)
	for i := 0; i < N; i++ {
		done[i] = false
	}

	for i := 0; i < N; i++ {
		if indeg[i] == 0 {
			rec(i, length, done, edges)
		}
	}
	fmt.Println(maxInt(length))
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
