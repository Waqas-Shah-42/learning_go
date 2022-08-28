package main

import (
	"fmt"
	"math"
)

func main() {
    fmt.Println("Hello, World!")
	temp := [5]int{1,2,3,4,5}

	for i, val := range(temp) {
		fmt.Printf("%v",i)
		fmt.Printf("%v",val)
	}

	fmt.Print(math.Pow(10,6))
	fmt.Printf("Miimum is %v",math.Min(temp)

	hello := average([4]int{4000, 3000, 1000, 2000})
}



func average(salary []int) float64 {

    sum := 0
    max := 0
    min := 1000000
    index := 0
    for i,val := range(salary) {
        index += 1
        sum += val
        min = int(math.Min(min,val))
        max = math.Max(max,val)
    }
    fmt.Printf("%v,%v,%v,%v", min,max,index,sum)
    
    return (sum-min-max)/float64(index-2)
}
