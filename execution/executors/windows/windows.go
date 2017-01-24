package windows

import (
	"context"
	"errors"
	"os"

	"github.com/docker/containerd/execution"
)

type WindowsRuntime struct{}

func New(ctx context.Context, root, shim, runtime string, runtimeArgs []string) (*WindowsRuntime, error) {
	return &WindowsRuntime{}, nil
}

func (s *WindowsRuntime) Create(ctx context.Context, id string, o execution.CreateOpts) (*execution.Container, error) {
	return nil, errors.New("create() not implemented")
}

func (s *WindowsRuntime) Delete(ctx context.Context, c *execution.Container) error {
	return errors.New("delete() not implemented")
}

func (s *WindowsRuntime) DeleteProcess(ctx context.Context, c *execution.Container, id string) error {
	return errors.New("deleteProcess() not implemented")
}

func (s *WindowsRuntime) List(ctx context.Context) ([]*execution.Container, error) {
	// Deliberately do NOT error to allow the Windows containerd to come-up to steady-state during initial bring-up
	return nil, nil
}

func (s *WindowsRuntime) Load(ctx context.Context, id string) (*execution.Container, error) {
	return nil, errors.New("load() not implemented")
}

func (s *WindowsRuntime) Pause(ctx context.Context, c *execution.Container) error {
	return errors.New("pause() not implemented")
}

func (s *WindowsRuntime) Resume(ctx context.Context, c *execution.Container) error {
	return errors.New("resume() not implemented")
}

func (s *WindowsRuntime) SignalProcess(ctx context.Context, c *execution.Container, id string, sig os.Signal) error {
	return errors.New("signalProcess() not implemented")
}

func (s *WindowsRuntime) Start(ctx context.Context, c *execution.Container) error {
	return errors.New("start() not implemented")
}

func (s *WindowsRuntime) StartProcess(ctx context.Context, c *execution.Container, o execution.StartProcessOpts) (p execution.Process, err error) {
	return nil, errors.New("startProcess() not implemented")
}
