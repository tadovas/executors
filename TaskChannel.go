package executors

import (
	"fmt"
	"sync"
)

type taskChannel chan Task

type taskChannelHolder struct {
	channel taskChannel
	wg      sync.WaitGroup
}

type errorCallback func(interface{})

func callSafely(task Task, handleError errorCallback) {
	defer func() {
		if val := recover(); val != nil {
			handleError(val)
		}
	}()
	task()
}

var DefaultCallback = func(panicVal interface{}) {
	fmt.Println("Last defence error:", panicVal)
}

func makeNewChannel(capacity int) (channel taskChannel) {
	channel = make(taskChannel, capacity)
	return
}

func consumeTasks(channel taskChannel, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if task := <-channel; task != nil {
			callSafely(task, DefaultCallback)
		} else {
			//shutdown on channel
			break
		}
	}
}

func (ste *taskChannelHolder) Execute(task Task) {
	ste.channel <- task
}

func (ste *taskChannelHolder) Shutdown() {
	close(ste.channel)
	ste.wg.Wait()
}
