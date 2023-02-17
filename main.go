/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/cemanaral/kron-cli/cmd"
	"github.com/cemanaral/kron-cli/pkg"
)

func main() {
	pkg.LoadHosts()
	cmd.Execute()
}
