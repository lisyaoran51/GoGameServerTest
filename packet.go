package main

import (
	"bytes"
	"encoding/binary"
	"math"
)

type Packet struct {
	Command  uint32
	Size     uint32
	Sequence uint32
	Body     []byte
}

//CmdKeepAlive KeepAlive
const CmdKeepAlive uint32 = 0x00000001

//SizeKeepAlive .
const SizeKeepAlive = 12

//PkKeepAlive KeepAlive
type PkKeepAlive struct {
	Head PacketHeader
}

//CmdCompress Compress
const CmdCompress uint32 = 0x00000002

//SizeCompress .
const SizeCompress = 18

//PkCompress Compress
type PkCompress struct {
	Head PacketHeader
	Type uint8  /** < \brief 壓縮方式. 0 - gzip, 1 - zlib, 2 - lzma */
	Len  uint32 /** < \brief 數據的原始大小 */
	//Data []byte//寫入時補上
}

//CmdRedirectIP RedirectIP
const CmdRedirectIP uint32 = math.MaxUint32 //equal to C++ -1

//SizeRedirectIP .
const SizeRedirectIP = 0

//PkRedirectIP RedirectIP
type PkRedirectIP struct {
	Head PacketHeader
	IP   uint32
}

//CmdProtoBuf ProtoBuf
const CmdProtoBuf uint32 = 0x00000008

//SizeProtoBuf .
const SizeProtoBuf = PacketHeaderSize + 4 + 1

//PkProtoBuf ProtoBuf
type PkProtoBuf struct {
	Packet    Packet
	ProtoSize uint32
	ProtoData []byte
}

//將 protobuf 的 data buffer 包裝成 PkProtoBuf
func ProtoToPackets(data []byte, size int) []byte {

	header := PacketHeader{
		Cmd:  CmdProtoBuf,
		Size: uint32(SizeProtoBuf - 1 + size),
	}

	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, header)
	binary.Write(buffer, binary.BigEndian, uint32(size))
	binary.Write(buffer, binary.BigEndian, data[:size])

	return buffer.Bytes()
}
