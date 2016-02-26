package supervisor

type StartTask struct {
	baseTask
	ID            string
	BundlePath    string
	Stdout        string
	Stderr        string
	Stdin         string
	StartResponse chan StartResponse
	Labels        []string
}

// Checkpoint not supported on Windows
func (task *startTask) setTaskCheckpoint(t *StartTask) {
}
