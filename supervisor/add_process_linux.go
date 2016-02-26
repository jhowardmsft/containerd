package supervisor

import (
	"github.com/opencontainers/specs"
)

type AddProcessTask struct {
	baseTask
	ID            string
	PID           string
	Stdout        string
	Stderr        string
	Stdin         string
	ProcessSpec   *specs.Process
	StartResponse chan StartResponse
}
