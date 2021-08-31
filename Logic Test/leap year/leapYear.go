package main

import "fmt"

func leapYear(a, b int){
	for a <= b{
		if(a % 4 == 0){
			fmt.Printf("%d ", a)
		}
		a++
	}
}

func main(){
	//leapYear(2000, 2020)
	leapYear(1900, 2020)
}