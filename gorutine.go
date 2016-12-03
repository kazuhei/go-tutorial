package main

import (
	"log"
	"strconv"
	"sync"
	"time"
)

func main() {
	wait := new(sync.WaitGroup)
	log.Printf("start")
	for i := 1; i < 5; i++ {
		wait.Add(1)
		go func(index int) {
			time.Sleep(time.Second * time.Duration(5-index))
			s := strconv.Itoa(index)
			log.Printf(s)
			wait.Done()
		}(i)
	}
	log.Printf("wait start")
	wait.Wait()
	log.Printf("end")
}
