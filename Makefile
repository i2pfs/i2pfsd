
protobufs:
	make -C protobuf
	make -C serverToServer/protobuf
	make -C serverToClient/protobuf

all: protobufs

