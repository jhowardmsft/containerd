package runtime

import (
	"io"
	"os"

	// This is a hack for Windows compilation
	"github.com/docker/containerd/windowsspecs"
)

// Duplicated from linx version to import the hacked version of specs.Process
type Process interface {
	io.Closer

	// ID of the process.
	// This is either "init" when it is the container's init process or
	// it is a user provided id for the process similar to the container id
	ID() string
	CloseStdin() error
	Resize(int, int) error
	// ExitFD returns the fd the provides an event when the process exits
	ExitFD() int
	// ExitStatus returns the exit status of the process or an error if it
	// has not exited
	ExitStatus() (int, error)
	// Spec returns the process spec that created the process
	Spec() specs.Process
	// Signal sends the provided signal to the process
	Signal(os.Signal) error
	// Container returns the container that the process belongs to
	Container() Container
	// Stdio of the container
	Stdio() Stdio
	// SystemPid is the pid on the system
	SystemPid() int
}

// Duplicated from linx version to import the hacked version of specs.Process
type process struct {
	root        string
	id          string
	pid         int
	exitPipe    *os.File
	controlPipe *os.File
	container   *container
	spec        specs.Process
	stdio       Stdio
}

// Duplicated from linx version to import the hacked version of specs.Process
type processConfig struct {
	id          string
	root        string
	processSpec specs.Process
	spec        *platformSpec
	c           *container
	stdio       Stdio
	exec        bool
	checkpoint  string
}

// TODO Windows: Linux uses syscalls which don't map to Windows. Needs alternate mechanism
func getExitPipe(path string) (*os.File, error) {
	return nil, nil
}

// TODO Windows: Linux uses syscalls which don't map to Windows. Needs alternate mechanism
func getControlPipe(path string) (*os.File, error) {
	return nil, nil
}

// TODO Windows. Windows does not support signals. Need alternate mechanism
// Signal sends the provided signal to the process
func (p *process) Signal(s os.Signal) error {
	return nil
}

func (p *process) Spec() specs.Process {
	return p.spec
}
func populateProcessStateForEncoding(config *processConfig, uid int, gid int) ProcessState {
	return ProcessState{
		Process: config.processSpec,
		Exec:    config.exec,
		Stdin:   config.stdio.Stdin,
		Stdout:  config.stdio.Stdout,
		Stderr:  config.stdio.Stderr,
	}
}
