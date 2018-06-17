package server

import (
	"github.com/i2pfs/i2pfsd/i2p"
	"github.com/i2pfs/i2pfsd/log"
	"github.com/i2pfs/i2pfsd/misc"
	pb "github.com/i2pfs/i2pfsd/protobuf/generated"
	cServer "github.com/i2pfs/i2pfsd/serverToClient/server"
	"github.com/i2pfs/i2pfsd/serverToServer/server"
	"github.com/majestrate/i2p-tools/sam3"
)

func helloMessageHandler(conn i2p.Connection, buf []byte) error {
	hello := pb.Hello{}
	err := hello.Unmarshal(buf)
	if err != nil {
		return err
	}
	log.Debugln("hello connection: received hello: ", hello)
	switch hello.ConnectionType {
	case pb.ConnectionType_Server:
		conn.SetMessageHandler(server.MessageHandler)
		err = conn.SendMessage(server.HelloMessage(conn, hello))
	case pb.ConnectionType_Client:
		conn.SetMessageHandler(cServer.MessageHandler)
		err = conn.SendMessage(cServer.HelloMessage(conn, hello))
	}
	return err
}

func handleListener(listener *sam3.StreamListener) {
	for {
		log.Debugln("Server Listener: Accept()...")
		rawConn, err := listener.Accept()
		misc.CheckError(err)
		log.Debugln("Server Listener: new connection")
		conn := i2p.NewConnection(rawConn)
		conn.SetMessageHandler(helloMessageHandler)
		conn.SendMessage(i2p.HelloMessage(conn))
		go conn.Loop()
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
	stream, err := sam.NewStreamSession("i2pfsd", keys, sam3.Options_Medium)
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
