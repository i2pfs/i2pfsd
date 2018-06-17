package main

import (
	"syscall"

	"github.com/xaionaro-go/log"
)

func setRLimitNoFile(newLimit uint64) error {
	var limit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &limit)
	if err != nil {
		log.Warningf("Got an error while syscall.Getrlimit(syscall.RLIMIT_NOFILE, &limit): %v", err)
		return err
	}
	if limit.Cur >= newLimit {
		return nil
	}
	limit.Cur = newLimit
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &limit)
	if err != nil {
		log.Warningf("Got an error while syscall.Setrlimit(syscall.RLIMIT_NOFILE, &limit): %v", err)
		return err
	}
	return nil
}
