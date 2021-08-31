package main

import "fmt"

func totalArray(a []int) int{
	i := 0
	for _, el := range a{
		i = i+el
	}
	return i
}

func nearestFib(a []int) int{
	tot := totalArray(a)
	startFib := []int{1,2}
	for !((startFib[0] <= tot) && (startFib[1] >= tot)) {
		nextFib := startFib[0] + startFib[1]
		startFib[0] = startFib[1]
		startFib[1] = nextFib
	}
	//fmt.Println(startFib)
	beforeDiff := tot - startFib[0]
	afterDiff := startFib[1] - tot
	if afterDiff <= beforeDiff{
		return afterDiff
	} else {
		return beforeDiff
	}
}

func main(){
	arr := []int{15,2,7,14}
	fmt.Println(nearestFib(arr))
}