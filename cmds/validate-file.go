package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	validateFileCmd.PersistentFlags().StringVarP(&inputPath, "input-path", "i", ".", "the directory to scan")
	validateFileCmd.PersistentFlags().IntVarP(&numWorkers, "workers", "w", 2, "the number of workers to allocate")
	ValidateCmd.AddCommand(validateFileCmd)
}

var validateFileCmd = &cobra.Command{
	Use: "file",
	Run: func(cmd *cobra.Command, args []string) {
		if info, err := os.Stat(inputPath); err == nil {
			if info.IsDir() {
				panic(fmt.Errorf("Use validate directory"))
			}
		} else {
			fmt.Fprintf(os.Stderr, err.Error()+"\n")
			os.Exit(1)
		}

		if err := validateFile(inputPath); err != nil {
			fmt.Printf("%s is invalid check log\n", inputPath)
			CloseLog()
			ScanLog()
		} else {
			fmt.Printf("%s is valid\n", inputPath)
			CloseLog()
		}

	},
}
