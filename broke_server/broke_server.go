package main

import (
	"bufio"
	"fmt"
	"os"
)
//7 1 4 1 9 1 1 9 1 7 9

func find(sl[]int) int{
	var f int
	var s int
	count := 0
	if len(sl) < 3 {
		return len(sl)
	}else{
		f = sl[0]
		for _, v := range sl{
			if f == v{
				count++
				continue
			}
			s = v
			count++
			break
		}
		r := count
		for i := count; i < len(sl); i++ {
			if sl[i] == f || sl[i] == s {
				r++
				if r > count{
					count = r
				}
				continue
			}
			if r > count{
				count = r
			}
			f = sl[i-1]
			r = 0
			for j := i - 1; j >= 0; j--{
				if sl[j] == f {
					r++
					continue
				}
				break
			}
			s = sl[i]
			r++
		}
		return count
	}

}
	

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int
	var count1 int
	var k int
	var res int
	fmt.Fscan(in, &count)
	for i := 0; i < count; i++{
		fmt.Fscan(in, &count1)
		sl := make([]int, count1)
		for j := 0; j < count1; j++{
			fmt.Fscan(in, &k)
			sl[j] = k
		}
		res = find(sl)
		fmt.Fprintln(out, res)
	}
}