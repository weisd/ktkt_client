package client

import (
	"github.com/hprose/hprose-go/hprose"
)

type ClientStub struct {
	RPCversion func() string `simple:"true"`
}

var ClientStubClient *ClientStub
