package main

import (
	"encoding/json"
	"os"
)

func main() {

	jsonString := os.Args[1]

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(jsonString), &jsonMap)

	a := jsonMap["a"].(int)
	b := jsonMap["b"].(int)

	println("Hello WasmEdge!")
	println("a 값은 아래 값입니다.")
	println(a)
	println("b 값은 아래 값입니다.")
	println(b)
	println(CustomFunc(int32(a + b)))
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
