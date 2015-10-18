package main

import (
	"fmt"
	"runtime"
)

func main() {
	// find the number of logical CPUs on the system
	totalCPUS := runtime.NumCPU()

	fmt.Println("total CPUs", totalCPUS)
	// set the golang runtime to use all the available processors
	runtime.GOMAXPROCS(totalCPUS)

	CPUWorkers := totalCPUS - 1
	fmt.Println("CPU workers", CPUWorkers)
}
