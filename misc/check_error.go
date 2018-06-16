package misc

import (
	"github.com/i2pfs/i2pfsd/log"
)

func CheckError(err error) {
	if err == nil {
		return
	}
	log.Panic(err)
}
