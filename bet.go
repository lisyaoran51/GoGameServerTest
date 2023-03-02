package main

import (
	"errors"
	"math/rand"
	"time"

	"github.com/lisyaoran51/GoGameServerTest/dao/clientDao"
	"github.com/lisyaoran51/GoGameServerTest/logger"
	"github.com/lisyaoran51/GoGameServerTest/packet"
	"github.com/lisyaoran51/GoGameServerTest/protobuf/flipCoin"
	"github.com/lisyaoran51/GoGameServerTest/protobuf/game"
	"github.com/shopspring/decimal"
)

func Bet(client *Client, reqHeader *flipCoin.BetReq) error {

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
		logger.Infof("[Bet] win %s\n", betAmount.Mul(decimal.NewFromInt(2)).String())
	} else {
		// lose
		logger.Infof("[Bet] lose %s\n", betAmount.String())
	}

	logger.Infof("[Bet] %s 發送扣款回應給玩家\n", clientModel.Username)
	//發送回應給client
	res := &game.GameMessage{
		Payload: &game.GameMessage_FlipCoinMessage{
			&flipCoin.GameMessage{
				Payload: &flipCoin.GameMessage_BetRes{
					&flipCoin.BetRes{
						Betamount: betAmount.String(),
						Code:      uint32(0),
						Bettime:   uint64(time.Now().Unix()),
						Win:       winAmount.String(),
					},
				},
			},
		},
	}

	resPacket := packet.NewProtobufPacket(0, res)

	connection := GetConnectionManager().GetConnection(GATE)
	if err = connection.SendPackage(client.SessionID, resPacket.ToByte()); err != nil {
		logger.Errorf("[Bet] %s 發送扣款回應給玩家失敗: %v", err)
		return err
	}

	return nil
}
