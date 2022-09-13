package main

import (
	"fmt"
	"sync"
)

var (
	counter1 int32
	counter2 int32
	wg       sync.WaitGroup
	mutex    sync.Mutex
	satu     interface{}
	dua      interface{}
)

func main() {
	wg.Add(2)

	go increment("satu")
	increment("dua")

	wg.Wait()
	fmt.Println("Counter:", counter1)

}

func increment(lang string) {

	satu = []string{"bisa1", "bisa2", "bisa3"}
	dua = []string{"coba1", "coba2", "coba3"}

	defer wg.Done()

	for i := 0; i < 8; i++ {
		mutex.Lock()
		{
			if lang == "satu" && i%2 == 0 {
				fmt.Println(satu, i+1)
				counter1++
			}
		}
		mutex.Unlock()

		if lang == "satu" && i%2 == 0 {
			fmt.Println(dua, i+1)
			counter2++
		}
	}
}
