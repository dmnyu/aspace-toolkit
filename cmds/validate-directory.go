package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/fs"
	"os"
	"path/filepath"
)

func init() {
	validateEadCmd.PersistentFlags().StringVarP(&inputPath, "input-path", "i", ".", "the directory to scan")
	validateEadCmd.PersistentFlags().IntVarP(&numWorkers, "workers", "w", 2, "the number of workers to allocate")
	ValidateCmd.AddCommand(validateEadCmd)
}

var (
	ignoreDirFilter = []string{".git", ".idea", ".circleci"}
	eadFiles        []string
)

func ignoreDir(name string) bool {
	for _, s := range ignoreDirFilter {
		if s == name {
			return true
		}
	}
	return false
}

var validateEadCmd = &cobra.Command{
	Use: "directory",
	Run: func(cmd *cobra.Command, args []string) {

		if info, err := os.Stat(inputPath); err == nil {
			//check if it is a file
			if !info.IsDir() {
				panic(fmt.Errorf("use validate file"))
			} else {
				if err = locateXMLFiles(); err != nil {
					(panic(err))
				}
				fmt.Printf("Gathered %d files\n", len(eadFiles))
				if err = validateFiles(); err != nil {
					panic(err)
				}
			}

		} else {
			fmt.Fprintf(os.Stderr, err.Error()+"\n")
			os.Exit(1)
		}
		CloseLog()
		ScanLog()
	},
}

func locateXMLFiles() error {
	err := filepath.WalkDir(inputPath, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			eadFiles = append(eadFiles, path)
		} else {
			if ignoreDir(d.Name()) {
				return filepath.SkipDir
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func validateFiles() error {
	workers := []int{}
	for i := 1; i <= numWorkers; i++ {
		workers = append(workers, i)
	}

	fmt.Printf("allocating %d workers\n", len(workers))
	resultChannel := make(chan int)
	for _, i := range workers {
		go allocateWorker(i, resultChannel)
	}

	count := 0
	for _ = range workers {
		count = count + <-resultChannel
	}

	fmt.Println(count, " ead files validated")

	return nil
}

func allocateWorker(id int, resultChannel chan int) {
	complete := false
	count := 0

	for complete != true {
		eadSlice := getEADFileSlice()
		if eadSlice == nil {
			complete = true
			break
		}
		for _, eadFile := range *eadSlice {
			validateFile(eadFile)
			count++
		}
	}
	resultChannel <- count
}

func getEADFileSlice() *[]string {
	if eadFiles == nil {
		return nil
	}
	if len(eadFiles) > 10 {
		eadSlice := eadFiles[:10]
		eadFiles = eadFiles[10:]
		return &eadSlice
	} else {
		eadSlice := eadFiles[:len(eadFiles)]
		eadFiles = nil
		return &eadSlice
	}
}
