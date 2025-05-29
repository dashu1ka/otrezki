package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var in *bufio.Reader
var out *bufio.Writer

type Otrezok struct {
	start int
	end   int
}

func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	otrezki := make([]Otrezok, n)
	for i := 0; i < n; i++ {
		var s, e int
		fmt.Fscan(in, &s, &e)
		if s > e {
			s, e = e, s
		}
		otrezki[i] = Otrezok{start: s, end: e}
	}

	fmt.Fprintln(out, solve(otrezki))

}

func solve(otrezki []Otrezok) int {
	sort.Slice(otrezki, func(i int, j int) bool {
		return otrezki[i].end < otrezki[j].end
	})
	count := 0
	point := -1
	for _, p := range otrezki {
		if point < p.start {
			point = p.end
			count++
		}
	}
	return count

}
