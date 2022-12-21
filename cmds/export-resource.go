package cmds

import (
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	exportResourceCmd.PersistentFlags().IntVar(&resource, "resource", 0, "")
	exportResourceCmd.PersistentFlags().IntVar(&repository, "repository", 0, "")
	exportResourceCmd.PersistentFlags().StringVarP(&config, "config", "c", "", "")
	exportResourceCmd.PersistentFlags().StringVarP(&environment, "environment", "e", "", "")
	exportResourceCmd.PersistentFlags().StringVarP(&outputDir, "outputDirectory", "o", ".", "")
	exportCmd.AddCommand(exportResourceCmd)
}

var exportResourceCmd = &cobra.Command{
	Use: "resource",
	Run: func(cmd *cobra.Command, args []string) {
		setClient()
		if err := exportResource(); err != nil {
			panic(err)
		}
	},
}

func exportResource() error {
	//get the ead as bytes
	eadBytes, err := client.GetEADAsByteArray(repository, resource, false)
	if err != nil {
		return err
	}

	asResource, err := client.GetResource(repository, resource)
	if err != nil {
		return err
	}

	//create the output filename
	eadFilename := strings.ToLower(asResource.MergeIDs("_")) + ".xml"
	eadFile := filepath.Join(outputDir, eadFilename)

	//write the file
	if err := os.WriteFile(eadFile, eadBytes, 0777); err != nil {
		return err
	}
	return nil
}
