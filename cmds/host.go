package cmds

import (
	"fmt"
	"github.com/nyudlts/go-aspace"
	"github.com/spf13/cobra"
)

func init() {
	hostCmd.PersistentFlags().StringVarP(&config, "config", "c", "", "")
	hostCmd.PersistentFlags().StringVarP(&environment, "environment", "e", "", "")
	rootCmd.AddCommand(hostCmd)
}

var hostCmd = &cobra.Command{
	Use: "host",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := aspace.NewClient(config, environment, 20)
		if err != nil {
			panic(err)
		}
		fmt.Println(client)
	}}
