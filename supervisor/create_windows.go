package supervisor

type platformStartTask struct {
}

// Checkpoint not supported on Windows
func setTaskCheckpoint(t *StartTask) {
}
