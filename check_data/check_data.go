package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int
	var count1 string
	sl := make([]string, 3)

	layoutISO := "2-1-2006"

	fmt.Fscan(in, &count)
	for i := 0; i < count; i++ {
		for j := 0; j < 3; j++ {
			fmt.Fscan(in, &count1)
			sl[j] = count1
		}
		str := strings.Join(sl, "-")
		_, err := time.Parse(layoutISO, str)
		if err != nil {
			fmt.Fprintln(out, "no")
		} else {
			fmt.Fprintln(out, "yes")
		}

	}

}
