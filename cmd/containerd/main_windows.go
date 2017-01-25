package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"unsafe"

	"google.golang.org/grpc"

	"github.com/Microsoft/go-winio"
	"github.com/Sirupsen/logrus"
	"github.com/docker/containerd/execution"
	"github.com/docker/containerd/execution/executors/windows"
	"github.com/docker/containerd/log"
	"github.com/docker/containerd/pkg/system"
	"github.com/urfave/cli"
)

const (
	listenerFlag         = "pipe"
	defaultRuntime       = "windows"
	numberSignalChannels = 2
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
	var (
		err      error
		executor execution.Executor
	)
	switch runtime {
	case "windows":
		executor, err = windows.New(log.WithModule(ctx, "windows"), root, "", "", nil)
		if err != nil {
			return nil, err
		}
		return executor, nil
	default:
		return nil, fmt.Errorf("runtime %q not implemented", runtime)
	}
}

func setupSignals(signals chan os.Signal) error {
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	// Windows does not support signals like *nix systems. So instead of
	// trapping on SIGUSR1 to dump stacks, we wait on a Win32 event to be
	// signaled. ACL'd to builtin administrators and local system
	ev := "Global\\containerd-" + fmt.Sprint(os.Getpid())
	sd, err := winio.SddlToSecurityDescriptor("D:P(A;;GA;;;BA)(A;;GA;;;SY)")
	if err != nil {
		return fmt.Errorf("failed to get security descriptor for debug stackdump event %s: %s", ev, err.Error())
	}
	var sa syscall.SecurityAttributes
	sa.Length = uint32(unsafe.Sizeof(sa))
	sa.InheritHandle = 1
	sa.SecurityDescriptor = uintptr(unsafe.Pointer(&sd[0]))
	h, err := system.CreateEvent(&sa, false, false, ev)
	if h == 0 || err != nil {
		return fmt.Errorf("failed to create debug stackdump event %s: %s", ev, err.Error())
	}
	go func() {
		logrus.Debugf("Stackdump - waiting signal at %s", ev)
		for {
			syscall.WaitForSingleObject(h, syscall.INFINITE)
			// TODO (@jhowardmsft) - Write to file, same as for docker/docker. Required otherwise
			// when running as service may blow the ETW event limit.
			dumpStacks()

			// TODO (@jhowardmsft): Similar to docker/docker - dump datastructures for debugging
			// path, err = d.dumpDaemon(root)
			// if err != nil {
			// logrus.WithError(err).Error("failed to write daemon datastructure dump")
			// } else {
			// logrus.Infof("daemon datastructure dump written to %s", path)
			// }
		}
	}()
	return nil
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
