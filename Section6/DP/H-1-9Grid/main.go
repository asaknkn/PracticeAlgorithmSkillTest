package main

import (
	"bufio"
	"fmt"
	"math"
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
	N := nextInt()
	M := nextInt()

	A := make([]string, N)

	for i := 0; i < N; i++ {
		A = append(A, next())
	}

	group := make([][]point, 11)

	for i := 0; i < N; i++ {
		//for j := 0; j < M; j++ {
		for j, c := range A[i] {
			fmt.Println(i, j, c)
			var n int
			if c == 'S' {
				n = 0
			} else if c == 'G' {
				n = 10
			} else {
				n = int(c) - 48
			}
			group[n] = append(group[n], point{x: i, y: j})
		}
	}
	fmt.Println(group)

	/*
		for i, v := range A {
			for j, c := range v {
				var n int
				if c == 'S' {
					n = 0
				} else if c == 'G' {
					n = 10
				} else {
					n = int(c) - 48
				}
				group[n] = append(group[n], point{x: i - 3, y: j})
			}
		}
	*/

	INF := math.MaxInt32
	cost := make([][]int, N)

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			cost[i] = append(cost[i], INF)
		}
	}

	si := group[0][0].x
	sj := group[0][0].y
	fmt.Println(group, cost)
	cost[si][sj] = 0

	for n := 1; n < 11; n++ {
		for _, v := range group[n] {
			i := v.x
			j := v.y

			for _, w := range group[n-1] {
				i2 := w.x
				j2 := w.y

				cost[i][j] = int(math.Min(float64(cost[i][j]), float64(cost[i2][j2])+math.Abs(float64((i-i2)+(j-j2)))))
			}
		}
	}

	gi := group[10][0].x
	gj := group[10][0].y
	ans := cost[gi][gj]
	if ans == INF {
		ans = -1
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
