package server

import (
	"bufio"
	"os"

	"github.com/i2pfs/i2pfsd/i2p"
	"github.com/i2pfs/i2pfsd/misc"
	pb "github.com/i2pfs/i2pfsd/protobuf/generated"
	cServer "github.com/i2pfs/i2pfsd/serverToClient/server"
	"github.com/i2pfs/i2pfsd/serverToServer/server"
	"github.com/majestrate/i2p-tools/sam3"
	"github.com/xaionaro-go/errors"
	"github.com/xaionaro-go/log"
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

func peersFileAddAddr(peersFilePath string, addr sam3.I2PAddr) error {
	addrString := addr.Base32()

	file, err := os.Open(peersFilePath)
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			peerAddress := scanner.Text()
			if peerAddress == addrString {
				file.Close()
				return nil
			}
		}
		file.Close()

		if err = scanner.Err(); err != nil {
			return errors.CannotParseFile.New(nil, "peersFileAddAddr")
		}
	}

	f, err := os.OpenFile(peersFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return errors.CannotOpenFile.New(err, "peersFileAddAddr[append]", peersFilePath)
	}
	defer f.Close()

	text := addrString+"\n"
	if _, err = f.WriteString(text); err != nil {
		return errors.CannotWriteToFile.New(err, "peersFileAddAddr", peersFilePath, text)
	}

	return nil
}

func Start(samUrl, peersFilePath, keysFilePath string) error {
	log.Debugln("Server: NewSAM")
	sam, err := sam3.NewSAM(samUrl)
	if err != nil {
		return errors.UnableToConnect.New(err, "i2p/server/server.go:Start():NewSAM()", samUrl)
	}

	log.Debugln("Server: EnsureKeyfile")
	keys, err := sam.EnsureKeyfile(keysFilePath)
	if err != nil {
		return errors.UnableToGetKey.New(err, "i2p/server/server.go:Start():EnsureKeyfile()", keysFilePath)
	}

	log.Debugln("Server: NewStreamSession")
	streamSession, err := sam.NewStreamSession("i2pfsd", keys, sam3.Options_Medium)
	if err != nil {
		return errors.UnableToStartSession.New(err, "i2p/server/server.go:Start():NewStreamSession()")
	}

	log.Debugln("Server: Adding me to the peers file if I'm not there")
	err = peersFileAddAddr(peersFilePath, streamSession.Addr())
	if err != nil {
		return errors.CannotWriteToFile.New(err, "i2p/server/server.go:Start():peersFileAddAddr()", peersFilePath, streamSession.Addr())
	}

	log.Debugln("Server: Listen")
	listener, err := streamSession.Listen()
	if err != nil {
		return errors.UnableToListen.New(err, "i2p/server/server.go:Start():Listen()")
	}

	go handleListener(listener)
	return nil
}
