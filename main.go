package main

import (
	"github.com/i2pfs/i2pfsd/config"
	"github.com/i2pfs/i2pfsd/i2p/server"
	"github.com/i2pfs/i2pfsd/misc"
)

func main() {
	cfg := config.GetDefaults()
	cfg.ReloadConfig()

	err := server.Start(cfg.SamUrl, cfg.KeysFilePath)
	misc.CheckError(err)

	for {
	}
}
