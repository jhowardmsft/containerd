package server

import (
	"errors"
	"fmt"

	"github.com/docker/containerd/api/grpc/types"
	"github.com/docker/containerd/supervisor"
	"github.com/docker/containerd/windowsspecs"
	"golang.org/x/net/context"
)

// noop on Windows (Checkpoints not supported)
func createContainerConfigCheckpoint(e *supervisor.StartTask, c *types.CreateContainerRequest) {
}

// TODO Windows - this can probably be factored better.
func (s *apiServer) AddProcess(ctx context.Context, r *types.AddProcessRequest) (*types.AddProcessResponse, error) {
	process := &specs.Process{
		Terminal: r.Terminal,
		Args:     r.Args,
		Env:      r.Env,
		Cwd:      r.Cwd,
	}
	if r.Id == "" {
		return nil, fmt.Errorf("container id cannot be empty")
	}
	if r.Pid == "" {
		return nil, fmt.Errorf("process id cannot be empty")
	}
	e := &supervisor.AddProcessTask{}
	e.ID = r.Id
	e.PID = r.Pid
	e.ProcessSpec = process
	e.Stdin = r.Stdin
	e.Stdout = r.Stdout
	e.Stderr = r.Stderr
	e.StartResponse = make(chan supervisor.StartResponse, 1)
	s.sv.SendTask(e)
	if err := <-e.ErrorCh(); err != nil {
		return nil, err
	}
	<-e.StartResponse
	return &types.AddProcessResponse{}, nil
}

// TODO Windows - may be able to completely factor out
func (s *apiServer) CreateCheckpoint(ctx context.Context, r *types.CreateCheckpointRequest) (*types.CreateCheckpointResponse, error) {
	return nil, errors.New("CreateCheckpoint() not supported on Windows")
}

// TODO Windows - may be able to completely factor out
func (s *apiServer) DeleteCheckpoint(ctx context.Context, r *types.DeleteCheckpointRequest) (*types.DeleteCheckpointResponse, error) {
	return nil, errors.New("DeleteCheckpoint() not supported on Windows")
}

// TODO Windows - may be able to completely factor out
func (s *apiServer) ListCheckpoint(ctx context.Context, r *types.ListCheckpointRequest) (*types.ListCheckpointResponse, error) {
	return nil, errors.New("ListCheckpoint() not supported on Windows")
}

func (s *apiServer) Stats(ctx context.Context, r *types.StatsRequest) (*types.StatsResponse, error) {
	return nil, errors.New("Stats() not supported on Windows")
}

func setUserFieldsInProcess(p *types.Process, oldProc specs.Process) {
}
