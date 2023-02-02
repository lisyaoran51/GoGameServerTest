package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/lisyaoran51/GoGameServerTest/protobuf"
	"github.com/lisyaoran51/GoGameServerTest/protobuf/diamonds"
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

		header := &protobuf.Header{}
		proto.Unmarshal(pk.ProtoData, header)

		if v.State == State_Connecting {

			switch header.Payload.(type) {
			case *protobuf.Header_GateLoginRequest:
				// TODO: add client
				header := header.GetGateLoginRequest()
				client := &Client{
					UserName: header.Name,
					IP:       v.IP,
				}
				v.Client = client

				// header := header.GetGateLoginRequest()
				// session.onGameClientConnect(header)

				// TODO: response
				// req := &protobuf.Header{
				// 	Uuid: protohelper.GenUUID(),
				// 	Payload: &protobuf.Header_GateLoginResponse{
				// 		&protobuf.GateLoginResponse{
				// 			Code:           uint32(code),
				// 			Username:       username,
				// 			CryptoCurrency: loginInfo.CurrentCurrency,
				// 			FiatCurrency:   loginInfo.CurrentFiatCurrency,
				// 		},
				// 	},
				// }

				// data := protohelper.MarshalHeader(req)

				// bufferData := packets.ProtoToPackets(data, len(data))
				// session.Write(bufferData, len(bufferData))
				// if code == errTable.ErrOk {
				// 	SendClientExchangeRate(obj, session)
				// 	onFinalSuccess(loginInfo, obj, session)
				// 	bufferData := GenerateCurrencyListProto(loginInfo.CurrencyData)
				// 	session.Write(bufferData, len(bufferData))
				// }
				return nil
			}
		}

		v.handleProtobuf(header)

		return true

	}
	return nil
}

func (v *VirtualSession) handleProtobuf(header *protobuf.Header) error {
	client := v.Client
	if client == nil {
		fmt.Printf("[Gate] session %#v 收到protobuf client是空，略過\n", v)
		return nil
	}
	//logger.Get().Infof("[Client] %s 收到protobuf %s", client.GetLoginname(), header.String())
	switch header.Payload.(type) {
	case *protobuf.Header_DiamondsCHeader:
		param := header.GetDiamondsCHeader()
		switch param.Payload.(type) {
		case *diamonds.CHeader_DiamondsTsBet:
			//請求下注
			reqHeader := param.GetDiamondsTsBet()
			err := Bet(client, reqHeader)

			if err != nil {
				//失敗了就提前回應，不然要等API扣款完成

				req := &protobuf.Header{
					Payload: &protobuf.Header_DiamondsCHeader{
						&diamonds.CHeader{
							Payload: &diamonds.CHeader_DiamondsBetRes{
								&diamonds.DIAMONDS_BET_RES{
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
