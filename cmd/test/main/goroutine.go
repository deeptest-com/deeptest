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
	for i := 1; i < 4; i++ {
		jobId := fmt.Sprintf("%d", i)

		ch := make(chan int, 10)

		jobChannelMap.Store(jobId, ch)

		go run(jobId, ch)
	}

	time.Sleep(time.Duration(10) * time.Second)
	stop("2")

	for i := 1; i < 1000000; i++ {
		time.Sleep(time.Duration(3) * time.Second)
		jobChannelMap.Range(func(key, value interface{}) bool {
			log.Println(key)
			log.Println(value)

			return true
		})
	}
}

func run(jobId string, ch chan int) {
	for i := 1; i < 1000000; i++ {
		// 耗时操作
		time.Sleep(time.Duration(3) * time.Second)

		select {
		case <-ch:
			goto GOTO

		default:
			log.Println(i)
		}
	}

GOTO:

	if ch != nil {
		jobChannelMap.Delete(jobId)
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
