package clientDao

import "errors"

var clients map[string]*ClientModel = make(map[string]*ClientModel)

type ClientModel struct {
	ID     string
	Amount string
}

func New(id string, amount string) error {
	clients[id] = &ClientModel{
		ID:     id,
		Amount: amount,
	}
	return nil
}

func Get(id string) *ClientModel {
	return clients[id]
}

func Modify(id string, amount string) error {
	if _, ok := clients[id]; !ok {
		return errors.New("no such client")
	}
	clients[id].Amount = amount
	return nil
}
