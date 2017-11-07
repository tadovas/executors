package executors

import "sync"

type threadPool struct {
	taskChannelHolder
}

func ThreadPool(numThreads int, channelSize int) Executor {
	wg := sync.WaitGroup{}
	wg.Add(numThreads)
	tp := threadPool{taskChannelHolder{make(taskChannel, channelSize), wg}}

	for i := 0; i < numThreads; i++ {
		go consumeTasks(tp.channel, &tp.wg)
	}

	return &tp
}
