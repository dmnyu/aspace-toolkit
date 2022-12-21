package cmds

import (
	"github.com/nyudlts/go-aspace"
	"github.com/spf13/cobra"
)

var (
	config      string
	environment string
	repository  int
	resource    int
	outputDir   string
	client      *aspace.ASClient
)

func init() {
	rootCmd.AddCommand(exportCmd)
}

func setClient() {
	var err error
	client, err = aspace.NewClient(config, environment, 20)
	if err != nil {
		panic(err)
	}
}

var exportCmd = &cobra.Command{
	Use: "export",
	Run: func(cmd *cobra.Command, args []string) {},
}
