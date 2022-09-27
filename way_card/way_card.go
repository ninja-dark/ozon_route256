package main

import (
	"bufio"
	"fmt"
	"os"
)
type field struct{
	row int
	col int
	r rune
}



func printWay( sl[][]string) {
	
	s := '*' 
	result := make([]field, 0)

	cache := make([][]rune, len(sl))
	for i, arr := range sl {
		cache[i] = make([]rune, len(arr))
	}
	queue := make([]field, 0)
	for k := range sl{
		for _, l := range sl[k]{
			str := []rune(l)
			for m, n := range str{
				queue = append(queue, field{k, m, n})
			}
		}
	}
	for ; len(queue) > 0 ; {
		p := queue[0]
		queue = queue[1:]
		if p.r == s {
			result = append(result, field{p.row, p.col, p.r})
		}
	} 
	
	fmt.Println(result)
}

func main(){
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var count int
	var a, b int
	var m string
	
	fmt.Fscan(in, &count)
	

	for i := 0; i < count; i++{
		for j := 0; j < count; j++{
			fmt.Fscan(in, &a, &b)
			matrix := make([][]string, 0)
			for j := 0; j < a; j++{
				fmt.Fscan(in, &m)
				temp := make([]string, 0)
				temp = append(temp, m)
				matrix = append(matrix, temp)
			}
			printWay(matrix)
		}
		
	}
}