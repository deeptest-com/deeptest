package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	jobChannelMap sync.Map
)

func main() {
	// start 3 jobs
	for i := 1; i <= 3; i++ {
		jobId := fmt.Sprintf("%d", i)

		ch := make(chan int, 10)

		jobChannelMap.Store(jobId, ch)

		go run(jobId, ch)
	}

	time.Sleep(time.Duration(5) * time.Second)
	// stop the 2nd job
	stop("2")
	time.Sleep(time.Duration(3) * time.Second)

	// check channel is nil and job removed
	jobChannelMap.Range(func(key, value interface{}) bool {
		log.Println(key)

		ch := value.(chan int)
		log.Println(ch)

		return true
	})
}

func run(jobId string, ch chan int) {
	for true {
		// 耗时操作
		time.Sleep(time.Duration(3) * time.Second)

		select {
		case <-ch:
			goto GOTO

		default:
		}
	}

GOTO:

	if ch != nil {
		//jobChannelMap.Delete(jobId)
		close(ch)
	}
}

func stop(jobId string) {
	chObj, ok := jobChannelMap.Load(jobId)

	if !ok || chObj == nil {
		return
	}

	ch := chObj.(chan int)
	if ch != nil {
		if !isChanClose(ch) {
			ch <- 1
		}

		ch = nil
	}
}

func isChanClose(ch chan int) bool {
	select {
	case <-ch:
		return true
	default:
	}
	return false
}
