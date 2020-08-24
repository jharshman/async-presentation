// not runnable OMIT
myJob := async.Job{
	Run: func() error {
		// execute some task(s)
	},
	Close: func() error {
		// gracefully shutdown task(s)
	},
  // specify what syscalls will trigger the Close() function
	Signals: []os.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL},

}

myJob.Execute()
