package supervisor

import "github.com/docker/containerd/runtime"

type platformStartTask struct {
	Checkpoint *runtime.Checkpoint
}

func setTaskCheckpoint(t *StartTask, task *startTask) {
	if t.Checkpoint != nil {
		task.Checkpoint = t.Checkpoint.Name
	}
}
