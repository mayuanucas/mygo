package main

import (
	"runtime"
	"sync"
	"fmt"
)

func main() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")
	go func() {
		defer wg.Done()

		for i := 0; i < 3; i++ {
			for char := 'a'; char < 'a'+26; char ++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 3; i++ {
			for char := 'A'; char < 'A'+26; char ++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Waiting To Finish")

	wg.Wait()
	fmt.Println("\nTerminating Program")
}
