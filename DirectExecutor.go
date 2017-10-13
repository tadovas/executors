package executors



type directExecutor struct {

}

func Direct() Executor {

	return directExecutor{}
}

func (de directExecutor ) Execute(task Task) {
	callSafely(task , DefaultCallback)
}

func (de directExecutor) Shutdown() {
	//nothing to do
}