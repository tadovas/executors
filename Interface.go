package executors

type Task func()

type Executor interface {
	Execute(Task)
	Shutdown()
}
