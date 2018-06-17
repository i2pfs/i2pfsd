package misc

import (
	"github.com/xaionaro-go/log"
)

func CheckError(err error) {
	if err == nil {
		return
	}
	log.Panic(err)
}
