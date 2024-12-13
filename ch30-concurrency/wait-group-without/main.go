package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Program start")
	for i := 0; i < 10; i++ {
		go concurrentTask(i)
	}
	finishTask()
	fmt.Println("Program end")
}

func finishTask() {
	fmt.Println("Executing finish task")
}

func concurrentTask(taskNumber int) {
	fmt.Printf("BEGIN Execute task number %d\n", taskNumber)
	time.Sleep(time.Millisecond * 100)
	fmt.Printf("END Execute task number %d\n", taskNumber)
}
