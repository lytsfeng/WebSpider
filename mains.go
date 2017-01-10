package main

import (
	"fmt"
	"time"
	//"sync"
	"sync"
)

var chan1 chan int
var chanLength int = 18
var interval time.Duration = 1500 * time.Millisecond

type poe struct {
	rwmutex sync.RWMutex
	name string

}

func (p * poe) Init(name string)  {
	p.rwmutex.Lock()
	defer p.rwmutex.Unlock()
	time.Sleep(time.Second * 1)
	p.name = "this..........."+name


}

func main() {

	p := &poe{}
	go func() {
		p.Init("one")
		fmt.Println(p)
	}()


	go func() {
		p.Init("two")
		fmt.Println(p)
	}()

	time.Sleep(time.Second* 15)


	//chan1 = make(chan int, chanLength)
	//
	//go func() {
	//	for i := 1; i < chanLength; i++ {
	//		if i > 0 && i%3 == 0 {
	//			fmt.Println("Reset chan1....")
	//			chan1 = make(chan int, chanLength)
	//		}
	//		fmt.Printf("send element %d \n", i)
	//		chan1 <- i
	//		time.Sleep(interval)
	//	}
	//	fmt.Println("Close chan1....")
	//	close(chan1)
	//}()
	//
	//receive()

}

func getChan() chan int {
	return chan1
}

func receive() {

	fmt.Println("begin to receive elements from chan1")
	timer := time.After(30 * time.Second)
Loop:
	for {
		select {
		case e, ok := <-getChan():
			if !ok {
				fmt.Println("--Closed chan1")
				break Loop
			}
			fmt.Printf("Received an element:%d\n", e)
			time.Sleep(interval)
		case <-timer:
			fmt.Println("timeout!")
			break Loop
		}
	}
	fmt.Println("-- End")

}
