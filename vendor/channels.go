package main

import "fmt"
func sum(list []int, c chan int)  {
	sum := 0
	for _,v :=range list{
		sum += v
	}

	// Send sum to chanel c
	c <- sum
}

func main() {
	//s := []int{7,2,8,-9,4,0}
	//
	//c := make(chan int)
	//go sum(s[len(s)/2:],c)
	//go sum(s[:len(s)/2],c)
	//
	//x,y := <-c,<-c // Receive from c
	//
	//fmt.Println(x,y)

	//Buffered Channels
	//If you send more than bufferred, it will deadlock
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
