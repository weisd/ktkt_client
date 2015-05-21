package client

type ClientStub struct {
	RPCversion func() string `simple:"true"`
}

var ClientStubClient *ClientStub
