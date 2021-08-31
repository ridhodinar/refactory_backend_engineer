package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(a int){
	i := 0
	res := make([]string, a)
	for i < a{
		if((i+1) % 15 == 0){
			res[i] = "FizzBuzz"
		} else if((i+1) % 5 == 0) {
			res[i] = "Buzz"
		} else if((i+1) % 3 == 0){
			res[i] = "Fizz"
		} else {
			res[i] = strconv.Itoa(i+1)
		}
		i++
	}
	fmt.Print(res)
}

func main(){
	fizzBuzz(15)
}