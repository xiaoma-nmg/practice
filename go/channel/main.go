package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var job chan *source
var result chan *out

type source struct {
	data int
}

type out struct {
	s   *source
	sum int
}

func producer(p chan<- *source) {
	defer wg.Done()
	for {
		s := &source{rand.Int()}
		p <- s
		time.Sleep(time.Millisecond * 500)
	}
}

func worker(p <-chan *source, r chan<- *out) {
	defer wg.Done()
	for {
		val := <-p
		n := val.data
		sum := 0
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		res := &out{
			s:   val,
			sum: sum,
		}
		r <- res
	}
}

func main() {
	job = make(chan *source, 100)
	result = make(chan *out, 100)
	wg.Add(1)
	go producer(job)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go worker(job, result)
	}

	for res := range result {
		fmt.Printf("source is %d, sum is %d\n", res.s.data, res.sum)
	}
	wg.Wait()
}
