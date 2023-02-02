package main

//PacketHeader 固定的封包標頭
type PacketHeader struct {
	Cmd  uint32
	Size uint32
	Seq  uint32
}

//PacketHeaderSize 封包標頭的大小
const PacketHeaderSize = 12

//CmdPackage .
const CmdPackage uint32 = 0x80000000

//PackageSize .
const PackageSize = 12

//PkPackage .
type PkPackage struct {
	Head PacketHeader
}

//CommandTypeMap 命令快查表
var CommandTypeMap = map[uint32]string{
	CmdPackage:          "CmdPackage",
	CmdPackageCompress:  "CmdPackageCompress",
	CmdClientEnter:      "CmdClientEnter",
	CmdClientEnterR:     "CmdClientEnterR",
	CmdClientExit:       "CmdClientExit",
	CmdClientExitR:      "CmdClientExitR",
	CmdGroupCreate:      "CmdGroupCreate",
	CmdGroupDestroy:     "CmdGroupDestroy",
	CmdGroupEnter:       "CmdGroupEnter",
	CmdGroupExit:        "CmdGroupExit",
	CmdGroupBroadcast:   "CmdGroupBroadcast",
	CmdGroupCache:       "CmdGroupCache",
	CmdClientRedirectIP: "CmdClientRedirectIP",
}

//CmdPackageCompress .
const CmdPackageCompress uint32 = 0x80000001

//PackageCompressSize 預先算好的PkPackageCompress的大小，對應C++原版的結構
const PackageCompressSize = 18

//PkPackageCompress .
type PkPackageCompress struct {
	Head         PacketHeader
	CompressType uint8
	OrgSize      uint32
	//Data         uint8
}

//ClientEnterSize .
const ClientEnterSize = PacketHeaderSize + 4

//CmdClientEnter .
const CmdClientEnter uint32 = 0x81000001

//PkClientEnter .
type PkClientEnter struct {
	Head PacketHeader
	IP   uint32
}

//CmdClientEnterR .
const CmdClientEnterR uint32 = 0x81001001

//ClientEnterRSize .
const ClientEnterRSize = PacketHeaderSize + 4

//PkClientEnterR .
type PkClientEnterR struct {
	Head PacketHeader
	Code uint32
}

//CmdClientRedirectIP .
const CmdClientRedirectIP uint32 = 0xffffffff

//ClientRedirectIPSize .
const ClientRedirectIPSize = PacketHeaderSize + 4

//PkClientRedirectIP .
type PkClientRedirectIP struct {
	Head PacketHeader
	IP   uint32
}

//ClientExitSize .
const ClientExitSize = PacketHeaderSize

//CmdClientExit .
const CmdClientExit uint32 = 0x81000002

//PkClientExit .
type PkClientExit struct {
	Head PacketHeader
}

//ClientExitRSize .
const ClientExitRSize = PacketHeaderSize + 4

//CmdClientExitR .
const CmdClientExitR uint32 = 0x81001002

//PkClientExitR .
type PkClientExitR struct {
	Head PacketHeader
	Code uint32
}

//GroupCreateSize .
const GroupCreateSize = PacketHeaderSize + 4

//CmdGroupCreate .
const CmdGroupCreate uint32 = 0x81003001

//PkGroupCreate .
type PkGroupCreate struct {
	Head    PacketHeader
	GroupID uint32
}

//GroupDestroySize .
const GroupDestroySize = PacketHeaderSize + 4

//CmdGroupDestroy .
const CmdGroupDestroy uint32 = 0x81003002

//PkGroupDestroy .
type PkGroupDestroy struct {
	Head    PacketHeader
	GroupID uint32
}

//GroupEnterSize .
const GroupEnterSize = PacketHeaderSize + 8

//CmdGroupEnter .
const CmdGroupEnter uint32 = 0x81003003

//PkGroupEnter .
type PkGroupEnter struct {
	Head     PacketHeader
	GroupID  uint32
	ClientID uint32
}

//GroupExitSize .
const GroupExitSize = PacketHeaderSize + 8

//CmdGroupExit .
const CmdGroupExit uint32 = 0x81003004

//PkGroupExit .
type PkGroupExit struct {
	Head     PacketHeader
	GroupID  uint32
	ClientID uint32
}

//GroupBroadcastSize .
const GroupBroadcastSize = PacketHeaderSize + 4

//CmdGroupBroadcast .
const CmdGroupBroadcast uint32 = 0x81003005

//PkGroupBroadcast .
type PkGroupBroadcast struct {
	Head    PacketHeader
	GroupID uint32
}

//CmdGroupCache .
const CmdGroupCache uint32 = 0x81003006

//GroupCacheSize 封包標頭的大小
const GroupCacheSize = 32

//PkGroupCache .
type PkGroupCache struct {
	Head    PacketHeader
	GroupID uint32
	CacheID uint32
	Flags   uint32
	/**< \brief 緩存旗標
	bit 0 = 1 時也廣播給當下在群組內的客戶
	bit 1 = 1 時 Append 在原有資料後面
	*/
	Expire int64
}
