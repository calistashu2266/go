package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		//wg sync.WaitGroup
		ch chan int = make(chan int, 20)
	)

	for index := 0; index < 20; index++ {
		//wg.Add(1)
		go func(i int) {
			//defer wg.Done()
			ch <- i
			//fmt.Println(i)
		}(index)
	}
	//wg.Wait()
	//close(ch)

	//for value := range ch {
	//	fmt.Printf("%d \n", value)
	//}
	ticker := time.NewTicker(time.Duration(3) * time.Second)

	for {

		select {
		case ps := <-ch:
			fmt.Printf("end:%d \r\n", ps)
		case time := <-time.After(time.Duration(2) * time.Second):
			fmt.Println("wait", time.Format("2006-01-02 15:04:05"))
		case times := <-ticker.C:
			fmt.Println("time", times)

		}
	}

	fmt.Println("end")
}
