package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"google.golang.org/grpc"

	"github.com/Microsoft/go-winio"
	"github.com/Sirupsen/logrus"
	"github.com/docker/containerd/execution"
	"github.com/urfave/cli"
)

const (
	listenerFlag   = "pipe"
	defaultRuntime = "windows"
)

var defaultRoot = filepath.Join(os.Getenv("PROGRAMDATA"), "containerd")

func appendPlatformFlags(flags []cli.Flag) []cli.Flag {
	return append(flags, cli.StringFlag{
		Name:  "pipe, p",
		Usage: "named pipe for containerd's GRPC server",
		Value: "//./pipe/containerd",
	})
	return flags
}

func processRuntime(ctx context.Context, runtime string, root string) (execution.Executor, error) {
	// TODO
	return nil, nil
}

func setupSignals(signals chan os.Signal) {
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
}

func handleSignals(signals chan os.Signal, server *grpc.Server) {
	for s := range signals {
		switch s {
		default:
			logrus.WithField("signal", s).Info("containerd: stopping GRPC server")
			server.Stop()
			return
		}
	}
}

func createListener(path string) (net.Listener, error) {
	// allow Administrators and SYSTEM
	sddl := "D:P(A;;GA;;;BA)(A;;GA;;;SY)"
	c := winio.PipeConfig{
		SecurityDescriptor: sddl,
		MessageMode:        true,  // Use message mode so that CloseWrite() is supported
		InputBufferSize:    65536, // Use 64KB buffers to improve performance
		OutputBufferSize:   65536,
	}
	l, err := winio.ListenPipe(path, &c)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on named pipe %q: %s. Is this process elevated?", path, err)
	}
	return l, nil
}
