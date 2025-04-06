package tasks

import (
	"fmt"
	"sync"
)

// Создайте две горутины, чтобы числа из одного канала читались по мере поступления,
// возводились в квадрат и результат записывался во второй канал.

func ReadSqrt() {
	chan1 := make(chan int)
	chan2 := make(chan int, 10)

	go func() {
		defer close(chan1)
		for i := range 10 {
			chan1 <- i
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range chan1 {
			chan2 <- v * v
			fmt.Println(v * v)
		}
	}()
	wg.Wait()
}
