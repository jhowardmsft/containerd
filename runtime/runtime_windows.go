package runtime

import "github.com/docker/containerd/windowsspecs"

// Checkpoint is not supported on Windows.
// TODO Windows: Can eventually be factored out entirely.
type Checkpoint struct {
}

// PlatformProcessState container platform-specific fields in the ProcessState structure
type PlatformProcessState struct {
}

type ProcessState struct {
	specs.Process
	Exec   bool   `json:"exec"`
	Stdin  string `json:"containerdStdin"`
	Stdout string `json:"containerdStdout"`
	Stderr string `json:"containerdStderr"`

	PlatformProcessState
}
