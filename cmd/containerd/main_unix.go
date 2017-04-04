package main

import (
	"os/signal"
	"syscall"

	"github.com/containerd/containerd/log"
	"github.com/containerd/containerd/reaper"
	"github.com/containerd/containerd/sys"
)

func handleSignals(signals chan os.Signal, server *grpc.Server) error {
	for s := range signals {
		log.G(global).WithField("signal", s).Debug("received signal")
		switch s {
		case syscall.SIGCHLD:
			if err := reaper.Reap(); err != nil {
				log.G(global).WithError(err).Error("reap containerd processes")
			}
		default:
			server.Stop()
			return nil
		}
	}
	return nil
}

func configureReaper() error {
	if conf.Subreaper {
		log.G(global).Info("setting subreaper...")
		if err := sys.SetSubreaper(1); err != nil {
			return err
		}
	}
}
