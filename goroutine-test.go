package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func tick(){
	tick := time.Tick(1000 * time.Millisecond)
	boom := time.After(5000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
func main() {
	//runtime.SetCPUProfileRate(500)
	cpuf, err := os.Create("cpu_profile")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(cpuf)
	defer pprof.StopCPUProfile()
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 30; i++ {
			fmt.Println(<-c)
			time.Sleep(500 * time.Millisecond)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}

