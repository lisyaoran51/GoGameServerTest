package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"net"
)

const (
	//MaxStackSize buffer緩衝大小
	MaxStackSize = 1024 * 64
)

func NewConnection(connection *net.Conn) *Connection {
	c := &Connection{
		Connection:      connection,
		VirtualSessions: make(map[uint32]*VirtualSession),
	}

	c.BufferReader = bufio.NewReaderSize(*connection, MaxStackSize)
	c.BufferWriter = bufio.NewWriterSize(*connection, MaxStackSize)

	return c
}

type Connection struct {
	Connection      *net.Conn
	BufferReader    *bufio.Reader
	BufferWriter    *bufio.Writer
	VirtualSessions map[uint32]*VirtualSession
}

func (c *Connection) GetVirtualSession(sequence uint32) *VirtualSession {
	if vs, ok := c.VirtualSessions[sequence]; ok {
		return vs
	}
	return nil
}

func (c *Connection) OnData(data []byte, length int) error {

	fmt.Printf("len: %d, recv: %q\n", length, string(data[:length]))

	dataBuffer := bytes.NewBuffer(data)

	// 讀標頭
	packet := &Packet{}

	binary.Read(dataBuffer, binary.BigEndian, &packet.Command)
	binary.Read(dataBuffer, binary.BigEndian, &packet.Size)
	binary.Read(dataBuffer, binary.BigEndian, &packet.Sequence)

	packet.Body = dataBuffer.Next(int(packet.Size) - PacketHeaderSize)

	fmt.Printf("packet: %+v\n", packet)

	c.OnPacket(packet)

	return nil
}

func (c *Connection) OnPacket(packet *Packet) error {

	buffer := bytes.NewBuffer(packet.Body)

	switch packet.Command {
	case CmdPackage:
		fmt.Printf("CmdPackage:\n")
		packageSize := int(packet.Size) - PacketHeaderSize
		packageData := buffer.Next(packageSize)
		if packet.Sequence == 0 {
			//gateway data
			c.OnData(packageData, packageSize)

		} else {
			if vs, ok := c.VirtualSessions[packet.Sequence]; ok {
				//gatesession_packethandling.go/onGameClientConnect

				dataBuffer := bytes.NewBuffer(packageData)
				customPacket := &Packet{}

				binary.Read(dataBuffer, binary.BigEndian, &customPacket.Command)
				binary.Read(dataBuffer, binary.BigEndian, &customPacket.Size)
				binary.Read(dataBuffer, binary.BigEndian, &customPacket.Sequence)

				customPacket.Body = dataBuffer.Next(int(customPacket.Size) - PacketHeaderSize)

				vs.OnPacket(customPacket)
			}
		}

		return nil
	case CmdClientEnter:
		fmt.Printf("CmdClientEnter:\n")

		if _, ok := c.VirtualSessions[packet.Sequence]; !ok {
			var clientIP uint32
			binary.Read(buffer, binary.BigEndian, &clientIP)
			vsession := NewVirtualSession(c, packet.Sequence, clientIP)
			c.VirtualSessions[packet.Sequence] = vsession
			fmt.Printf("virtual session added to [%d]\n", packet.Sequence)
		}
		return nil
	}
	return nil
}

func (c *Connection) SendPackage(seq int, body []byte, size int) error {

	pkPackage := PkPackage{}

	pkPackage.Head.Cmd = CmdPackage
	pkPackage.Head.Seq = uint32(seq)
	pkPackage.Head.Size = uint32(PackageSize + size)

	//tmpData := pbytes.GetCap(gatecommand.PackageSize + size)
	tmpData := make([]byte, 0, PackageSize+uint32(PackageSize+size))
	buffer := bytes.NewBuffer(tmpData)

	binary.Write(buffer, binary.BigEndian, pkPackage)

	buffer.Write(body[:size])

	_, err := c.BufferWriter.Write(buffer.Bytes())

	if err != nil {
		return err
	}
	err = c.BufferWriter.Flush()
	if err != nil {
		return err
	}
	return nil
}

func (c *Connection) Start(ctx context.Context) {

	fmt.Printf("Connection開始連線讀取\n")

	message := make([]byte, MaxStackSize)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Session停止連線讀取\n")
			return
		default:
		}

		//stcpsession.conn.SetReadDeadline(time.Now().Add(config.SocketTimeout * time.Second))
		//message := pbytes.GetLen(MaxStackSize)
		length, err := c.BufferReader.Read(message) //read 是 blocking  的操作，所以可以用在for loop中
		//fmt.Println("CTCPSession puller leng:", length)
		if err != nil {
			fmt.Printf("連線錯誤 %v\n", err)
			return
		}
		if length > 0 {
			c.OnData(message[:length], length)
		}
	}
}
