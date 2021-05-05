package main

import (
	"bufio"
	"fmt"
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
	h := nextInt()
	w := nextInt()

	s := make([][]rune, h)
	for i := 0; i < h; i++ {
		ss := next()
		for _, v := range ss {
			s[i] = append(s[i], v)
		}
	}

	var si, sj, gi, gj int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == 's' {
				si = i
				sj = j
			}
			if s[i][j] == 'g' {
				gi = i
				gj = j
			}
		}
	}

	visited := make([][]bool, h)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			visited[i] = append(visited[i], false)
		}
	}

	dfs(h, w, si, sj, visited, s)

	if visited[gi][gj] == true {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

func dfs(h, w, i, j int, visited [][]bool, s [][]rune) {
	visited[i][j] = true

	xdirection := []int{0, 1, 0, -1}
	ydirection := []int{1, 0, -1, 0}

	for k := 0; k < 4; k++ {
		i2 := i + xdirection[k]
		j2 := j + ydirection[k]

		if (0 > i2 || i2 >= h) || (0 > j2 || j2 >= w) {
			continue
		}

		if s[i][j] == '#' {
			continue
		}

		if visited[i2][j2] != true {
			dfs(h, w, i2, j2, visited, s)
		}
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
