package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		fmt.Println("begin")
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}



//ch <- v    // Send v to channel ch.
//v := <-ch  // Receive from ch, and
//// assign value to v.


func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		// This method will wait till fibonacci run
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
