package main

import (
	"os"
	"strconv"
)

func main() {

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	println("Hello WasmEdge!")
	println(CustomFunc(int32(n)))
}

func CustomFunc(n int32) int32 {
	arr := make([]int32, n)
	for i := int32(0); i < n; i++ {
		switch {
		case i < 2:
			arr[i] = i
		default:
			arr[i] = arr[i-1] + arr[i-2]
		}
	}
	return arr[n-1]
}
