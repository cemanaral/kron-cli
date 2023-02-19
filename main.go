/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/cemanaral/kron/cmd"
	"github.com/cemanaral/kron/pkg/host"
)

func main() {
	host.LoadHosts()
	cmd.Execute()
}
