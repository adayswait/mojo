package msignal

import (
	"github.com/adayswait/mojo/mlog"
	"os"
	"os/signal"
	"syscall"
)

var signch = make(chan os.Signal, 1)

func init() {
	signal.Notify(signch,
		syscall.SIGABRT,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM)
}

func CatchSignal() {
	for {
		s := <-signch
		switch s {
		case syscall.SIGABRT:
			fallthrough
		case syscall.SIGHUP:
			fallthrough
		case syscall.SIGINT:
			fallthrough
		case syscall.SIGQUIT:
			fallthrough
		case syscall.SIGTERM:
			mlog.Infof("caught signal %s, exit", s.String())
			mlog.Flush()
			os.Exit(0)
		default:
			mlog.Infof("caught signal %s, nop", s.String())
			mlog.Flush()
		}
	}
}
