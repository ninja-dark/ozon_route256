package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

//2 1 3 1 1 4
//0 1 2 3 4 5
//l m
//
func sum1(sl[]int, lim int) {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	 minIndex := 1
	 n := sl[0]
	 res := make([]int, len(sl))
	 res = append(res, n)
	 count:= lim
	
	minValue := math.Abs(float64(n) - float64(sl[1]))
	if count > 0 {
		count--
		for num := 2; num < len(sl); num++ {
			if math.Abs((float64(n) - float64(sl[num]))) < minValue{
				res = append(res, sl[num])
				minIndex = num
				minValue = math.Abs((float64(n) - float64(sl[num])))
			}
			n = minIndex
			res = append(res, sl[num])
			
		}
	}
	fmt.Println(res)

	for k := range res[5:]{
		fmt.Fprintln(out, k+1 )
	}
	

}


func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int
	var count1 int
	var s int
	
	
//	tems:= []int{}
	fmt.Fscan(in, &count)
	for i := 0; i < count; i++{
		fmt.Fscan(in, &count1)
		sl := make([]int, count1)
		for j := 0; j < count1; j ++ {
			fmt.Fscan(in, &s)
			sl[j] = s
		}
		sum1(sl, count1)
		
		
	}
}