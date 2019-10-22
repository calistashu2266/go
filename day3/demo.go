package main

import (
	"fmt"
	"sync"
)

type person struct {
	age  int
	name string
}

func (p *person) Say() {
	fmt.Printf("age=%d,name=%s \r\n", p.age, p.name)
}

func main() {
	var p2 *person = new(person)
	p2.name = "dddd"
	p2.age = 444
	p2.Say()

	var p1 *person = &person{24, "ggggg"}
	p1.Say()

	var ch chan int = make(chan int, 100)
	// var wg sync.WaitGroup
	var wg *sync.WaitGroup = new(sync.WaitGroup)
	for index := 0; index < 100; index++ {
		wg.Add(1)
		go func(index int) {
			ch <- index
			defer wg.Done()
		}(index)
	}
	wg.Wait()
	for {
		select {
		case x := <-ch:
			fmt.Println(x)
		}
	}

	// defer close(ch)

}
