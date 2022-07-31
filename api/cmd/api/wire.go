//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

func InitHttpProtocol() interface{} {
	wire.Build()
	return nil
}
