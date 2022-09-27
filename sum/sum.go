package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int 
	fmt.Fscan(in, &count)

	for i := 0; i <count; i++ {
		var n, m int 
		fmt.Fscan(in, &n, &m)
		fmt.Fprintln(out, n + m)
	}


}