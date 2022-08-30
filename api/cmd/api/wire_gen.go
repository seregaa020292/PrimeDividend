// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"primedivident/internal/infrastructures/server/http"
)

// Injectors from wire.go:

func InitializeServer() http.Server {
	server := http.NewServer()
	return server
}
