package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	reindexCmd.PersistentFlags().StringVarP(&config, "config", "c", "", "")
	reindexCmd.PersistentFlags().StringVarP(&environment, "environment", "e", "", "")
	rootCmd.AddCommand(reindexCmd)
}

var reindexCmd = &cobra.Command{
	Use: "reindex",
	Run: func(cmd *cobra.Command, args []string) {
		setClient()
		code, err := client.Reindex()
		if err != nil {
			panic(err)
		}
		fmt.Println(code)

	},
}
