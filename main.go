package main

import (
	example5fanoutworkers "concurrency/example5_fan_out_workers"
	"fmt"
)

func main() {
	fmt.Println("Hello concurrency !!!")
	//example1.SomeFunc()
	//example2.SomeFunc()
	//example3.SomeFunc()
	//example4_fan_in.SomeFunc()
	example5fanoutworkers.SomeFunc()
}
