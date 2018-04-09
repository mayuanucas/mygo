package main

import (
	"fmt"
	"time"
	"github.com/mayuanucas/mygo/lib"
	"runtime"
)

func main() {
	//runtime.GOMAXPROCS(1)
	//
	//var wg sync.WaitGroup
	//wg.Add(2)
	//
	//fmt.Println("Start Goroutines")
	//go func() {
	//	defer wg.Done()
	//
	//	for i := 0; i < 3; i++ {
	//		for char := 'a'; char < 'a'+26; char ++ {
	//			fmt.Printf("%c ", char)
	//		}
	//	}
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//
	//	for i := 0; i < 3; i++ {
	//		for char := 'A'; char < 'A'+26; char ++ {
	//			fmt.Printf("%c ", char)
	//		}
	//	}
	//}()
	//
	//fmt.Println("Waiting To Finish")
	//
	//wg.Wait()
	//fmt.Println("\nTerminating Program")

	pool := lib.New(runtime.NumCPU() + 1)
	for i := 1; i <= 5; i++ {
		go lowwerCase(i, pool)
		go upperCase(i, pool)
	}
	pool.WaitAll()

	fmt.Println("all done.")
}

func lowwerCase(id int, group *lib.Pool) {
	group.Add(1)
	defer group.Done()

	fmt.Println("lowwerCase ID-->", id)
	for char := 'a'; char <= 'z'; char += 1 {
		fmt.Printf("%c ", char)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()
	fmt.Println("done-->", id)
}

func upperCase(id int, group *lib.Pool) {
	group.Add(1)
	defer group.Done()

	fmt.Println("upperCase ID-->", id)
	for char := 'A'; char <= 'Z'; char += 1 {
		fmt.Printf("%c ", char)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println()
	fmt.Println("done-->", id)
}
