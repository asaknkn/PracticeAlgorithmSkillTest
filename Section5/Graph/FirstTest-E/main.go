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

	n := nextInt()
	q := nextInt()

	graph := make([][]bool, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			graph[i] = append(graph[i], false)
		}
	}

	query := make([][]int, q)
	for i := 0; i < q; i++ {
		f := nextInt()
		query[i] = append(query[i], f)
		a := nextInt()
		query[i] = append(query[i], a)

		if f == 1 {
			b := nextInt()
			query[i] = append(query[i], b)
		}
	}

	for i := 0; i < q; i++ {
		a := query[i][1] - 1

		if query[i][0] == 1 {
			b := query[i][2] - 1
			graph[a][b] = true
		}

		if query[i][0] == 2 {
			for j := 0; j < n; j++ {
				if graph[j][a] == true {
					graph[a][j] = true
				}
			}
		}

		if query[i][0] == 3 {
			var toFollow []int
			for v := 0; v < n; v++ {
				if graph[a][v] == true {
					for w := 0; w < n; w++ {
						if graph[v][w] == true && w != a {
							toFollow = append(toFollow, w)
						}
					}
				}
			}

			for _, w := range toFollow {
				graph[a][w] = true
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if graph[i][j] == true {
				fmt.Printf("Y")
			} else {
				fmt.Printf("N")
			}
		}
		fmt.Printf("\n")
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
