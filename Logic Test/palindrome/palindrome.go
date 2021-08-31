package main

import (
	"fmt"
	"strings"
)

func palindrome(s string) string{
	lower := strings.ToLower(s)
	i := 0
	for i <= len(s)/2{
		if lower[i] != lower[len(s)-1-i]{
			return "not palindrome"
		}
		i++
	}
	return "palindrome"
}

func main(){
	fmt.Println(palindrome("Radar"))
	fmt.Println(palindrome("Malam"))
	fmt.Println(palindrome("Kasur ini rusak"))
	fmt.Println(palindrome("Ibu Ratna antar ubi"))
	fmt.Println(palindrome("Malas"))
	fmt.Println(palindrome("Makan nasi goreng"))
	fmt.Println(palindrome("Balonku ada lima"))
}