package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/lisyaoran51/GoGameServerTest/dao/clientDao"
	"github.com/lisyaoran51/GoGameServerTest/protobuf/flipCoin"
	"github.com/lisyaoran51/GoGameServerTest/protobuf/game"
	"google.golang.org/protobuf/proto"
)

func NewVirtualSession(c *Connection, clientID uint32, ip uint32) *VirtualSession {
	return &VirtualSession{
		Connection: c,
		ClientID:   clientID,
		IP:         ip,
	}
}

type VirtualSession struct {
	Connection *Connection
	ClientID   uint32
	IP         uint32
	State      State
	Client     *Client
}

func (v *VirtualSession) OnPacket(packet *Packet) error {
	fmt.Printf("OnPacket %+v\n", packet)
	if v.State == State_Disconnected {
		return errors.New("Disconnected")
	}

	cmd := packet.Command
	if cmd == CmdProtoBuf {
		if v.State != State_Connecting && v.State != State_Connected {
			fmt.Printf("wrong state %d\n", v.State)
			return errors.New("state error")
		}

		buffer := bytes.NewBuffer(packet.Body)
		pk := PkProtoBuf{
			Packet: *packet,
		}
		binary.Read(buffer, binary.BigEndian, &pk.ProtoSize)
		pk.ProtoData = buffer.Next(int(pk.ProtoSize))

		header := &game.GameMessage{}
		proto.Unmarshal(pk.ProtoData, header)
		fmt.Printf("get proto: %+v\n", header)

		if v.State == State_Connecting {

			switch header.Payload.(type) {
			case *game.GameMessage_LoginReq:
				// TODO: add client
				header := header.GetLoginReq()
				client := &Client{
					UserName:  header.Name,
					IP:        v.IP,
					SessionID: packet.Sequence,
				}
				v.Client = client

				clientDao.New(header.Name, "10000")

				req := &game.GameMessage{
					Payload: &game.GameMessage_LoginRes{
						&game.LoginRes{
							Code:     uint32(0),
							Username: header.Name,
						},
					},
				}

				dataBuffer, err := proto.Marshal(req)
				if err != nil {
					return err
				}

				bufferData := ProtoToPackets(dataBuffer, len(dataBuffer))

				if err = v.Connection.SendPackage(int(v.ClientID), bufferData, len(bufferData)); err != nil {
					return err
				}

				return nil
			}
		}

		v.handleProtobuf(header)

		return nil

	}
	return nil
}

func (v *VirtualSession) handleProtobuf(header *game.GameMessage) error {
	fmt.Printf("handleProtobuf %+v\n", header)
	client := v.Client
	if client == nil {
		fmt.Printf("[Gate] session %#v 收到protobuf client是空，略過\n", v)
		return nil
	}
	//logger.Get().Infof("[Client] %s 收到protobuf %s", client.GetLoginname(), header.String())
	switch header.Payload.(type) {
	case *game.GameMessage_FlipCoinMessage:
		param := header.GetFlipCoinMessage()
		switch param.Payload.(type) {
		case *flipCoin.GameMessage_BetReq:
			//請求下注
			reqHeader := param.GetBetReq()
			err := Bet(client, reqHeader)

			if err != nil {
				//失敗了就提前回應，不然要等API扣款完成

				req := &game.GameMessage{
					Payload: &game.GameMessage_FlipCoinMessage{
						&flipCoin.GameMessage{
							Payload: &flipCoin.GameMessage_BetRes{
								&flipCoin.BetRes{
									Code: uint32(111),
								},
							},
						},
					},
				}

				dataBuffer, err := proto.Marshal(req)
				if err != nil {
					return err
				}

				bufferData := ProtoToPackets(dataBuffer, len(dataBuffer))

				if err = v.Connection.SendPackage(int(v.ClientID), bufferData, len(bufferData)); err != nil {
					return err
				}
			}
		}
		return nil
	}
	return nil
}

func (v *VirtualSession) Write(body []byte, size int) error {

	return v.Connection.SendPackage(int(v.ClientID), body, size)

}
