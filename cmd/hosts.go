/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/cemanaral/kron/pkg/host"
	"github.com/spf13/cobra"
)

// hostsCmd represents the hosts command
var hostsCmd = &cobra.Command{
	Use:   "hosts",
	Short: "Prints loaded host info",
	Long: "Prints loaded host information from /etc/kron/hosts.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		printHosts(cmd, args)
	},
}

func printHosts(cmd *cobra.Command, args []string) {
	fmt.Printf("Total number of hosts loaded: %d\n\n", len(host.Hosts))
	for _, host := range host.Hosts {
		fmt.Println(host)
	}
}

func init() {
	rootCmd.AddCommand(hostsCmd)
}
