package executors

import "sync"

type singleThread struct {
	taskChannelHolder
}

func Serial(channelSize int) Executor {
	wg := sync.WaitGroup{}
	wg.Add(1)
	st := singleThread{taskChannelHolder{makeNewChannel(channelSize), wg}}
	go consumeTasks(st.channel, &st.wg)
	return &st
}
