package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

type Connection struct {
	BufferReader    *bufio.Reader
	BufferWriter    *bufio.Writer
	VirtualSessions map[uint32]*VirtualSession
}

func (c *Connection) OnData(data []byte, length int) error {

	fmt.Printf("len: %d, recv: %s\n", length, string(data[:length]))

	dataBuffer := bytes.NewBuffer(data)

	// 讀標頭
	packet := &Packet{}

	binary.Read(dataBuffer, binary.BigEndian, &packet.Command)
	binary.Read(dataBuffer, binary.BigEndian, &packet.Size)
	binary.Read(dataBuffer, binary.BigEndian, &packet.Sequence)

	packet.Body = dataBuffer.Next(int(packet.Size) - PacketHeaderSize)

	c.OnPacket(packet)

	return nil
}

func (c *Connection) OnPacket(packet *Packet) error {

	buffer := bytes.NewBuffer(packet.Body)

	switch packet.Command {
	case CmdPackage:
		packageSize := int(packet.Size) - PacketHeaderSize
		packageData := buffer.Next(packageSize)
		if packet.Sequence == 0 {
			//gateway data
			c.OnData(packageData, packageSize)

		} else {
			if vs, ok := c.VirtualSessions[packet.Sequence]; ok {
				//gatesession_packethandling.go/onGameClientConnect
				vs.OnPacket(packageData, packageSize)
			}
		}

		return nil
	case CmdClientEnter:
		//fmt.Println(" gatecommand.CmdClientEnter clientID:", clientID)

		if _, ok := c.VirtualSessions[packet.Sequence]; !ok {
			var clientIP uint32
			binary.Read(buffer, binary.BigEndian, &clientIP)
			vsession := NewVirtualSession(c, packet.Sequence, clientIP)
			c.VirtualSessions[packet.Sequence] = vsession
		}
		return nil
	}

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
