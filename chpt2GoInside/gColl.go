package main

import (
	"fmt"
	"runtime"
	"time"
)

func printStats(mem runtime.MemStats)  {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC", mem.NextGC)
	fmt.Println("-------")
}

func main()  {
	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++{
		s := make([]byte, 5000000)
		if s == nil {
			fmt.Println("Operation Failf")
		}
	}
	printStats(mem)

	for i := 0; i < 10; i++{
		s := make([]byte, 15000000)
		if s == nil {
			fmt.Println("Operation Failf")
		}
		time.Sleep(time.Second)
	}
	printStats(mem)
}