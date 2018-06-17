package server

import (
	"net"

	"github.com/i2pfs/i2pfsd/log"
	"github.com/i2pfs/i2pfsd/misc"
	"github.com/i2pfs/i2pfsd/serverToServer/protobuf/generated"
	"github.com/majestrate/i2p-tools/sam3"
)

func handleConnection(connection net.Conn) {
	for {
		buf := make([]byte, 4096)
		log.Debugln("Server Connection: Read")
		n, err := connection.Read(buf)
		misc.CheckError(err)
		log.Debugln("Server Connection: received: " + string(buf[:n]))
		message := s2sProtobuf.Message{}
		err = message.Unmarshal(buf[:n])
		misc.CheckError(err)
		log.Debugln("Server Connection: received message: ", message)
	}
}

func handleListener(listener *sam3.StreamListener) {
	for {
		log.Debugln("Server Listener: Accept()...")
		conn, err := listener.Accept()
		misc.CheckError(err)
		log.Debugln("Server Listener: new connection")
		go handleConnection(conn)
	}
}

func Start(samUrl, keysFilePath string) error {
	log.Debugln("Server: NewSAM")
	sam, err := sam3.NewSAM(samUrl)
	if err != nil {
		return err
	}
	log.Debugln("Server: EnsureKeyfile")
	keys, err := sam.EnsureKeyfile(keysFilePath)
	if err != nil {
		return err
	}
	log.Debugln("Server: NewStreamSession")
	stream, err := sam.NewStreamSession("i2pfsd-s2s-s", keys, sam3.Options_Medium)
	if err != nil {
		return err
	}
	log.Debugln("Server: Listen")
	listener, err := stream.Listen()
	if err != nil {
		return err
	}
	go handleListener(listener)
	return nil
}
