package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/lisyaoran51/GoGameServerTest/dao/clientDao"
	"github.com/lisyaoran51/GoGameServerTest/protobuf/diamonds"
	"github.com/shopspring/decimal"
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

	if result == 1 {
		// win
		balance.Add(betAmount.Mul(decimal.NewFromInt(2)))
		clientDao.Modify(client.UserName, balance.String())
		fmt.Printf("win %s\n", betAmount.Mul(decimal.NewFromInt(2)).String())
	} else {
		// lose
		fmt.Printf("lose %s\n", betAmount.String())
	}

	return nil
}
