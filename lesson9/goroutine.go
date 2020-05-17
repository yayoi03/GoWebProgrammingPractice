// package main

// import (
// 	"fmt"
// 	"time"
// )

// // func printNumers1() {
// // 	for i := 0; i < 10; i++ {
// // 		//fmt.Printf("%d", i)
// // 	}
// // }

// // func printLetters1() {
// // 	for i := 'A'; i < 'A'+10; i++ {
// // 		//fmt.Printf("%c", i)
// // 	}
// // }

// // func print1() {
// // 	printNumers1()
// // 	printLetters1()
// // }

// // func goPrint1() {
// // 	go printNumers1()
// // 	go printLetters1()
// // }

// // func printNumers2(wg *sync.WaitGroup) {
// // 	for i := 0; i < 10; i++ {
// // 		time.Sleep(1 * time.Millisecond)
// // 		fmt.Printf("%d", i)
// // 	}
// // 	wg.Done()
// // }

// // func printLetters2(wg *sync.WaitGroup) {
// // 	for i := 'A'; i < 'A'+10; i++ {
// // 		time.Sleep(1 * time.Millisecond)
// // 		fmt.Printf("%c", i)
// // 	}
// // 	wg.Done()
// // }

// // func print2() {
// // 	printNumers2()
// // 	printLetters2()
// // }

// // func goPrint2() {
// // 	go printNumers2()
// // 	go printLetters2()
// // }

// func printNumers2(w chan bool) {
// 	for i := 0; i < 10; i++ {
// 		time.Sleep(1 * time.Millisecond)
// 		fmt.Printf("%d", i)
// 	}
// 	w <- true
// }

// func printLetters2(w chan bool) {
// 	for i := 'A'; i < 'A'+10; i++ {
// 		time.Sleep(1 * time.Millisecond)
// 		fmt.Printf("%c", i)
// 	}
// 	w <- true
// }

// func thrower(c chan int) {
// 	for i := 0; i < 5; i++ {
// 		c <- i
// 		fmt.Println("Threw >>", i)
// 	}
// }

// func catcher(c chan int) {
// 	for i := 0; i < 5; i++ {
// 		num := <-c
// 		fmt.Println("Caught <<", num)
// 	}
// }

// func callerA(c chan string) {
// 	c <- "Hello World!"
// 	close(c)
// }

// func callerB(c chan string) {
// 	c <- "Hola Mundo!"
// 	close(c)
// }

// func main() {
// 	// w1, w2 := make(chan bool), make(chan bool)
// 	// go printNumers2(w1)
// 	// go printLetters2(w2)
// 	// <-w1
// 	// <-w2

// 	// c := make(chan int, 3)
// 	// go thrower(c)
// 	// go catcher(c)
// 	// time.Sleep(300 * time.Microsecond)

// 	a, b := make(chan string), make(chan string)
// 	go callerA(a)
// 	go callerB(b)
// 	var msg string
// 	ok1, ok2 := true, true

// 	for ok1 || ok2 {
// 		select {
// 		case msg, ok1 = <-a:
// 			if ok1 {
// 				fmt.Printf("%s from A\n", msg)
// 			}
// 		case msg, ok2 = <-b:
// 			if ok2 {
// 				fmt.Printf("%s from B\n", msg)
// 			}
// 		}
// 	}
// }
