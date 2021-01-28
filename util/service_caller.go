package util

import "github.com/micro/go-micro/client"

var clients = client.DefaultClient

func InitClient(input client.Client){
	clients = input
}

func GetClient()client.Client{
	return clients
}