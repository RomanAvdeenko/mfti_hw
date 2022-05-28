package main

import (
	"fmt"
)

// сюда писать код
func ExecutePipeline(jobs MyStr) {
	//in := make(chan interface{})
	//out := make(chan interface{})

	//jobs(in, out)
	print(jobs)
}

func SingleHash(in, out chan interface{})     {}
func MultiHash(in, out chan interface{})      {}
func CombineResults(in, out chan interface{}) {}

func main() {
	ExecutePipeline("Test")
	fmt.Println(DataSignerMd5("0"))
	fmt.Println(DataSignerCrc32("0"))
}
