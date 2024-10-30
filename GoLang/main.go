package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	fmt.Println("Loading Programs......")

	//Run your imported modules here - Check ref.txt
	// Example -> resolve(integertoroman.Solve, 450)
}

// func resolve(algo interface{}, args ...interface{}) {
func resolve(algo interface{}, args ...interface{}) {
	start := time.Now()

	// Use reflection to call the function with arguments
	fn := reflect.ValueOf(algo)
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		fmt.Println(i, arg)
		in[i] = reflect.ValueOf(arg)
	}
	fn.Call(in)

	duration := time.Since(start)
	fmt.Printf("Algorithm took: %vms\n", duration.Milliseconds())
}
