package main

import (
	"fmt"
	"sync"
)

func NReinas(reinas [8]int, fila int) {
	it := 1
	n := 8
	if fila == n {
		var t [8][8]int
		for i := 0; i < 8; i++ {
			t[i] = [8]int{0, 0, 0, 0, 0, 0, 0, 0}
		}
		for x, y := range reinas {
			if y != -1 {
				t[x][y] = 1
			}
		}
		for _, val := range t {
			fmt.Println(val)
			it = it + 1

		}
	} else {
		for c := 0; c < n; c++ {
			reinas[fila] = c
			NReinas(reinas, fila+1)

		}

	}
}

func main() {
	row := [8]int{-1, -1, -1, -1, -1, -1, -1, -1}
	NReinas(row, 0)
}
