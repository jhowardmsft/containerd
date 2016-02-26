package supervisor

type StartTask struct {
	baseTask
	ID            string
	BundlePath    string
	Stdout        string
	Stderr        string
	Stdin         string
	StartResponse chan StartResponse
	Checkpoint    *runtime.Checkpoint
	Labels        []string
}

func setTaskCheckpoint(t *StartTask) {
	if t.Checkpoint != nil {
		task.Checkpoint = t.Checkpoint.Name
	}
}
