package main

import (
	"github.com/i2pfs/i2pfsd/config"
	"github.com/i2pfs/i2pfsd/i2p/server"
	"github.com/i2pfs/i2pfsd/misc"
	"github.com/i2pfs/i2pfsd/consts"
)

func main() {
	var err error

	cfg := config.GetDefaults()
	cfg.ReloadConfig()

	err = setRLimitNoFile(consts.RLIMIT_NOFILE)
	misc.CheckError(err)

	err = server.Start(cfg.SamUrl, cfg.PeersFilePath, cfg.KeysFilePath)
	misc.CheckError(err)

	for {
	}
}
