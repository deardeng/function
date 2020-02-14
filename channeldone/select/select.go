package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func makeIntChannel() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(15)) * time.Microsecond)
			out <- i
			i++
		}
	}()
	return out
}

func doWork(id int, w worker) {
	for n := range w.in {
		//time.Sleep(500 * time.Millisecond)
		fmt.Printf("Worker %d received %d\n", id, n)
	}
	w.done()
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func main() {
	c1, c2 := makeIntChannel(), makeIntChannel()
	var wg sync.WaitGroup
	var worker = createWorker(0, &wg)
	n := 0

	var values []int
	tm := time.After(100 * time.Second)
	tick := time.Tick(time.Second)

	wg.Add(1)

	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker.in
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tm:
			fmt.Println("bye")
			return
		case <-tick:
			fmt.Println("len of queue = ", len(values))
		}
	}
	wg.Done()

}
