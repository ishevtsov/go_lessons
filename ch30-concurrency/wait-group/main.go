package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Program start")
	// initialize the wait group
	var waitGroup sync.WaitGroup
	waitGroup.Add(10)
	for i := 0; i < 10; i++ {
		go concurrentTask(i, &waitGroup)
	}
	waitGroup.Wait()
	finishTask()
	fmt.Println("Program end")
}

func finishTask() {
	fmt.Println("Executing finish task")
}

func concurrentTask(taskNumber int, waitGroup *sync.WaitGroup) {
	fmt.Printf("BEGIN Execute task number %d\n", taskNumber)
	time.Sleep(time.Millisecond * 100)
	fmt.Printf("END Execute task number %d\n", taskNumber)
	waitGroup.Done()
}
