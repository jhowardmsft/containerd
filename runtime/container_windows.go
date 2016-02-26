package runtime

import (
	"errors"

	"github.com/docker/containerd/windowsspecs"
)

type Container interface {
	// ID returns the container ID
	ID() string
	// Path returns the path to the bundle
	Path() string
	// Start starts the init process of the container
	Start(checkpoint string, s Stdio) (Process, error)
	// Exec starts another process in an existing container
	Exec(string, specs.Process, Stdio) (Process, error)
	// Delete removes the container's state and any resources
	Delete() error
	// Processes returns all the containers processes that have been added
	Processes() ([]Process, error)
	// State returns the containers runtime state
	State() State
	// Resume resumes a paused container
	Resume() error
	// Pause pauses a running container
	Pause() error
	// RemoveProcess removes the specified process from the container
	RemoveProcess(string) error
	// Labels are user provided labels for the container
	Labels() []string
	// Pids returns all pids inside the container
	Pids() ([]int, error)
	// Stats returns realtime container stats and resource information
	Stats() (*Stat, error)
}

func getRootIDs(s *platformSpec) (int, int, error) {
	return 0, 0, nil
}

func (c *container) Pause() error {
	return errors.New("Pause not supported on Windows")
}

func (c *container) Resume() error {
	return errors.New("Resume not supported on Windows")
}

func (c *container) Checkpoints() ([]Checkpoint, error) {
	return nil, errors.New("Checkpoints not supported on Windows ")
}

func (c *container) Checkpoint(cpt Checkpoint) error {
	return errors.New("Checkpoint not supported on Windows ")
}

func (c *container) DeleteCheckpoint(name string) error {
	return errors.New("DeleteCheckpoint not supported on Windows ")
}

// TODO Windows: Implement me.
// This will have a very different implementation on Windows.
func (c *container) Start(checkpoint string, s Stdio) (Process, error) {
	return nil, errors.New("Start not yet implemented on Windows")
}

// TODO Windows: Implement me.
// This will have a very different implementation on Windows.
func (c *container) Exec(pid string, spec specs.Process, s Stdio) (Process, error) {
	return nil, errors.New("Exec not yet implemented on Windows")
}

// TODO Windows: Implement me.
func (c *container) Pids() ([]int, error) {
	return nil, errors.New("Pids not yet implemented on Windows")
}

// TODO Windows: Implement me. (Not yet supported by docker on Windows either...)
func (c *container) Stats() (*Stat, error) {
	return nil, errors.New("Stats not yet implemented on Windows")
}
