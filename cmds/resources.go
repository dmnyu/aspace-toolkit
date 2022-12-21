package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	resourceListCmd.PersistentFlags().IntVar(&repository, "repository", 0, "")
	resourceListCmd.PersistentFlags().StringVarP(&config, "config", "c", "", "")
	resourceListCmd.PersistentFlags().StringVarP(&environment, "environment", "e", "", "")
	rootCmd.AddCommand(resourceListCmd)
}

var resourceListCmd = &cobra.Command{
	Use: "resources",
	Run: func(cmd *cobra.Command, args []string) {
		setClient()
		resources, err := client.GetResourceIDs(repository)
		if err != nil {
			panic(err)
		}

		fmt.Println(resources)
	},
}
