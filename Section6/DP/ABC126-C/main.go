package main

import (
	"bufio"
	"fmt"
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

func dfs(i int, child [][]int) int {
	if len(child[i]) == 0 {
		//fmt.Println(i)
		return 1
	} else {
		var value []int
		for _, v := range child[i] {
			value = append(value, dfs(v, child))
		}
		return maxInt(value) + minInt(value) + 1
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	N := nextInt()
	child := make([][]int, N)
	//for i := 0; i < N; i++ {
	//	child[i] = make([]int, N)
	//}
	for i := 1; i < N; i++ {
		b := nextInt()
		child[b-1] = append(child[b-1], i)
	}

	//fmt.Println(child)
	fmt.Println(dfs(0, child))
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
