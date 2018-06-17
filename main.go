package main

import (
	"github.com/i2pfs/i2pfsd/config"
	"github.com/i2pfs/i2pfsd/misc"
	"github.com/i2pfs/i2pfsd/serverToServer/server"
	//"github.com/i2pfs/i2pfsd/serverToServer/client"
	//cServer "github.com/i2pfs/i2pfsd/serverToClient/server"
)

func main() {
	cfg := config.GetDefaults()
	cfg.ReloadConfig()

	err := server.Start(cfg.SamUrl, cfg.KeysFilePath)
	misc.CheckError(err)
}
