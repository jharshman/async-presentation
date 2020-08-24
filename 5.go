// not runnable OMIT
myJob := async.Job{
	Run: func() error {
		// execute some task(s)
	},
	Close: func() error {
		// gracefully shutdown task(s)
	},
}

myJob.Execute()
