package main

import (
        "fmt"
        "math/rand"
        "time"
)

func main() {

        //Static array
		//x := []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17,}
        
		//Create auto rand digits array (from 0 to 99) with lenth = count 
		fmt.Printf("Input count of digits in array (max 100)")
        var count int
        fmt.Scanf("%d", &count)
        rand.Seed(time.Now().UnixNano())
        x := make([]int, count)
        for i := 0; i < count; i++ {
                x[i] = rand.Intn(100)
        }
        fmt.Println(x)
		
		//Find MIN digit in array
        var min int = 100
        for _, i := range x {
                if min > i {
                        min = i
                }
        }
        fmt.Println("%v", min)
}

