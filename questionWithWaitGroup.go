package main

import (
	"fmt"
	"sync"
	"time"
)

/*
будут выводиться числа "у каждой горутины 1 число", "у  каждой горутины n+1 число" числа под индексом n В случайном порядке из 3х горутин
*/
// это то решение которое ожидалось? я книгу читал ( приостановил на изучение алгоритмов ), там как раз следующая тема была многопоточность ("язык программирования go" называется книга) если что продолжу\доизучу что нужно
func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		// если Done не выполнится из-за ошибки, будет deadlock, а defer выполняется при любом завершении функции
		defer wg.Done()
		for _, value := range []int{1, 4, 7} {
			fmt.Println(value)
			time.Sleep(time.Millisecond * 1)
		}
	}()

	go func() {
		defer wg.Done()
		for _, value := range []int{2, 5, 8} {
			fmt.Println(value)
			time.Sleep(time.Millisecond * 1)
		}
	}()

	go func() {
		defer wg.Done()
		for _, value := range []int{3, 6, 9} {
			fmt.Println(value)
			time.Sleep(time.Millisecond * 1)
		}
	}()
	wg.Wait()
}
