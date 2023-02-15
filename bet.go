package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/lisyaoran51/GoGameServerTest/dao/clientDao"
	"github.com/lisyaoran51/GoGameServerTest/protobuf"
	"github.com/lisyaoran51/GoGameServerTest/protobuf/diamonds"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/proto"
)

func Bet(client *Client, reqHeader *diamonds.DIAMONDS_TS_BET) error {

	clientModel := clientDao.Get(client.UserName)

	betAmount, err := decimal.NewFromString(reqHeader.Betamount)
	if err != nil {
		return err
	}

	balance, err := decimal.NewFromString(clientModel.Amount)
	if err != nil {
		return err
	}

	if balance.LessThan(betAmount) {
		return errors.New("balance not enough")
	}

	balance = balance.Sub(betAmount)
	clientDao.Modify(client.UserName, balance.String())

	rand.Seed(time.Now().UnixNano())

	result := rand.Int() % 2
	winAmount := decimal.Zero

	if result == 1 {
		// win
		winAmount = betAmount.Mul(decimal.NewFromInt(2))
		balance.Add(winAmount)
		clientDao.Modify(client.UserName, balance.String())
		fmt.Printf("win %s\n", betAmount.Mul(decimal.NewFromInt(2)).String())
	} else {
		// lose
		fmt.Printf("lose %s\n", betAmount.String())
	}

	fmt.Printf("[Client] %s 發送扣款回應給玩家\n", clientModel.Username)
	//發送回應給client
	cheader := &protobuf.Header_DiamondsCHeader{
		&diamonds.CHeader{
			Payload: &diamonds.CHeader_DiamondsBetRes{
				&diamonds.DIAMONDS_BET_RES{
					Betamount: betAmount.String(),
					Code:      uint32(0),
					Bettime:   uint64(time.Now().Unix()),
					Win:       winAmount.String(),
				},
			},
		}}

	dataBuffer, err := proto.Marshal(&protobuf.Header{
		Payload: cheader,
	})
	if err != nil {
		return err
	}

	bufferData := ProtoToPackets(dataBuffer, len(dataBuffer))

	connection := GetConnectionManager().GetConnection(GATE)
	if err = connection.SendPackage(int(client.SessionID), bufferData, len(bufferData)); err != nil {
		return err
	}

	return nil
}
