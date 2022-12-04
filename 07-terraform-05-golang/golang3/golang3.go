package main

import "fmt"

func main() {
	var x []int
	k :=3
	for i:=0; i<100; i+=k {
		x=append(x,i)
	}
	/*
	for i:=0; i<100; i++ {
		if i%k == 0 {
			x=append(x,i)
		}
	}
	*/
	fmt.Println(x)
}