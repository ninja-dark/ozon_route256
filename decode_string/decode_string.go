package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



func main(){
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	
	var count int
	var str string
	var res string
	var r string
	

	fmt.Fscan(in, &count)
	for i := 0; i < count; i++{
		fmt.Fscan(in, &str)
		sl := make([]string, len(str))
		
		sl = strings.Split(str, "")
		

		for _, v := range sl{
			res += v
			if res == "00"{
				r += "a"
				res = ""
			}
			if res == "100" {
				r += "b"
				res = ""
			}
			if res == "101"{
				r += "c"
				res = ""
			}
			if res == "11"{
				r += "d"
				res = ""
			}

			
		}
		fmt.Fprintln(out, r)
		r = ""
		
	}
	
}