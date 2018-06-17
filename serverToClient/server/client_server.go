package clientServer

import (
	"fmt"

	"github.com/i2pfs/i2pfsd/i2p"
	"github.com/i2pfs/i2pfsd/log"
	"github.com/i2pfs/i2pfsd/protobuf/generated"
	"github.com/i2pfs/i2pfsd/serverToClient/protobuf/generated"
)

func HelloMessage(connection i2p.Connection, hello helloProtobuf.Hello) *s2cProtobuf.Hello {
	return &s2cProtobuf.Hello{}
}

func MessageHandler(connection i2p.Connection, buf []byte) error {
	message := s2cProtobuf.Message{}
	err := message.Unmarshal(buf)
	if err != nil {
		return err
	}
	log.Debugln("cServer Connection: received message: ", message)
	return fmt.Errorf("Not implemented, yet")
}
