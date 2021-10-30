package main

import "tritonhttp"
import "fmt"

func main() {
	a, err := tritonhttp.ParseMIME("./mime.types")
	if err != nil {
		panic("fuck")
	}
	fmt.Println(a)
}
