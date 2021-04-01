package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	chan1 := make(chan struct{})
	chan2 := make(chan struct{})
	chan3 := make(chan struct{})

	go func() {
		defer wg.Done()
		for _, value := range []int{0, 3, 6, 9} {
			fmt.Println(value)
			chan1 <- struct{}{}
			<-chan3
		}
	}()

	go func() {
		defer wg.Done()
		for _, value := range []int{1, 4, 7, 10} {
			<-chan1
			fmt.Println(value)
			chan2 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for _, value := range []int{2, 5, 8, 11} {
			<-chan2
			fmt.Println(value)
			chan3 <- struct{}{}
		}
	}()
	wg.Wait()
}
