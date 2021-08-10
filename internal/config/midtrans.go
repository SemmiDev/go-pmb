package config

import (
	"github.com/veritrans/go-midtrans"
)

func NewMidtransClient(configuration Config) midtrans.Client {
	midClient := midtrans.NewClient()
	midClient.ServerKey = configuration.Get("MIDTRANS_SERVER_KEY")
	midClient.ClientKey = configuration.Get("MIDTRANS_CLIENT_KEY")
	midClient.APIEnvType = midtrans.Sandbox
	return midClient
}
