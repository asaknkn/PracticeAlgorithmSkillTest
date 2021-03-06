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
	c := "abcdefghijklmnopqrstuvwxyz."
	s := next()

	var cnt int
	for _, t := range chars(c) {
		if isMatch(t, s) {
			cnt++
		}
	}

	for _, c1 := range chars(c) {
		for _, c2 := range chars(c) {
			t := c1 + c2
			if isMatch(t, s) {
				cnt++
			}
		}
	}

	for _, c1 := range chars(c) {
		for _, c2 := range chars(c) {
			for _, c3 := range chars(c) {
				t := c1 + c2 + c3
				if isMatch(t, s) {
					cnt++
				}
			}
		}
	}

	fmt.Println(cnt)
}

func isMatch(t string, s string) bool {
	for i := 0; i < len(s)-len(t)+1; i++ {
		ok := true

		for j := 0; j < len(t); j++ {
			if s[i+j] != t[j] && t[j] != '.' {
				ok = false
			}
		}
		if ok {
			return true
		}
	}
	return false
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
