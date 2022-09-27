package main

import (
	"bufio"
	"fmt"
	"os"
)

func sum(s []int) int{
	m := make(map[int]int)
	sum := 0
	for _, v := range s{
		if _, ok := m[v]; ok {
			m[v]++
		}else {
			m[v] = 1
		}
	}
	for k,v := range m {
		h := v /3
		sum = sum + k *(v - h)
	}
	return sum

}

func main(){

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int
	var count1 int
	var s int
	fmt.Fscan(in, &count)
	for i := 0; i < count; i++{
		fmt.Fscan(in, &count1)
		sl := make([]int, count1)
		for j := 0; j < count1; j ++ {
			fmt.Fscan(in, &s)
			sl[j] = s
		}

		fmt.Fprintln(out, sum(sl))
	
	}
	
}