package server

import (
	"fmt"

	"github.com/i2pfs/i2pfsd/i2p"
	"github.com/xaionaro-go/log"
	"github.com/i2pfs/i2pfsd/protobuf/generated"
	"github.com/i2pfs/i2pfsd/serverToServer/protobuf/generated"
)

func HelloMessage(connection i2p.Connection, hello helloProtobuf.Hello) *s2sProtobuf.Hello {
	return &s2sProtobuf.Hello{}
}

func MessageHandler(connection i2p.Connection, buf []byte) error {
	message := s2sProtobuf.Message{}
	err := message.Unmarshal(buf)
	if err != nil {
		return err
	}
	log.Debugln("cServer Connection: received message: ", message)
	return fmt.Errorf("Not implemented, yet")
}
