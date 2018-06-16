package server

import (
	"github.com/i2pfs/i2pfsd/log"
	"github.com/i2pfs/i2pfsd/misc"
	"github.com/majestrate/i2p-tools/sam3"
)

func Run(samUrl, keysFilePath string) {
	log.Debugln("Server: NewSAM")
	sam, err := sam3.NewSAM(samUrl)
	misc.CheckError(err)
	log.Debugln("Server: NewKeys")
	keys, err := sam.NewKeys()
	misc.CheckError(err)
	log.Debugln("Server: NewStreamSession")
	stream, err := sam.NewStreamSession("serverTun1", keys, sam3.Options_Medium)
	misc.CheckError(err)
	log.Debugln("Server: Listen")
	listener, err := stream.Listen()
	misc.CheckError(err)
	log.Debugln("Server: Accept")
	conn, err := listener.Accept()
	misc.CheckError(err)
	buf := make([]byte, 4096)
	log.Debugln("Server: Read")
	n, err := conn.Read(buf)
	log.Debugln("Server received: " + string(buf[:n]))
}
