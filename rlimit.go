package main

import (
	"syscall"

	"github.com/xaionaro-go/log"
)

func setRLimitNoFile(newLimit uint64) error {
	var limit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &limit)
	if err != nil {
		return log.WarningWrapper(nil, err, "Got an error while syscall.Getrlimit(syscall.RLIMIT_NOFILE, &limit)")
	}
	if limit.Cur >= newLimit {
		return nil
	}
	limit.Cur = newLimit
	oldMax := limit.Max
	if limit.Max < limit.Cur {
		limit.Max = limit.Cur
	}
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &limit)
	if err == syscall.EPERM {
		if oldMax < limit.Max {
			log.Warningf("No permission to increase the limit of open files (nofile)")
			limit.Max = oldMax
			limit.Cur = oldMax
			err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &limit)
		}
	}
	if err != nil {
		return log.WarningWrapper(nil, err, "Got an error while syscall.Setrlimit(syscall.RLIMIT_NOFILE, &limit)")
	}
	return nil
}
