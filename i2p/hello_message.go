package i2p

import (
	pb "github.com/i2pfs/i2pfsd/protobuf/generated"
)

func HelloMessage(conn Connection) *pb.Hello {
	return &pb.Hello{}
}
