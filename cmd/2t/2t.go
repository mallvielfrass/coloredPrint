package main

import "fmt"

type Test struct {
	Data string
}
type Vall []Test

func main() {
	structTest := Test{
		Data: "test",
	}
	var a Vall
	a = append(a, Test{
		Data: "2",
	})
	a = append(a, structTest)
	fmt.Printf("%+v\n", a)
}
