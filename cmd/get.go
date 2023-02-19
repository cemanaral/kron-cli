package cmd

import (
	"log"

	"github.com/cemanaral/kron/pkg/host"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve cron informations from hosts",
	Long: `Retrieve cron informations from hosts`,
	Run: func(cmd *cobra.Command, args []string) {
		get(args)
	},
}

func get(args []string) {
	selectedHosts := make(host.HostList, len(args), len(args))
	for i, hostName := range args {
		hostIndex := slices.IndexFunc(host.Hosts, func(h host.Host) bool {return h.Id == hostName}); 
		if hostIndex == -1 {
			log.Fatalf("Host '%s' could not be found in %s. Please validate your configuration by executing 'kron hosts'", hostName, host.HOSTS_PATH)
		}
		selectedHosts[i] = host.Hosts[hostIndex] 
	}
	for _, h := range selectedHosts {
		h.GetCronInformation()
	}	
}

func init() {
	rootCmd.AddCommand(getCmd)
}
