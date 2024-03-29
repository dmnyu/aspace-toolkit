package cmds

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	inputPath  string
	logger     *os.File
	xmlPtn     = regexp.MustCompile("xml$")
	rootCmd    = &cobra.Command{}
	logFile    = "ead-tools.log"
	numWorkers = 8
)

func init() {
	var err error
	logger, err = os.Create(logFile)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logger)
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

func CloseLog() {
	logger.Close()
}

func ScanLog() {
	fmt.Println("** errors and warnings **")
	lf, err := os.Open(logFile)
	if err != nil {
		panic(err)
	}
	defer lf.Close()
	scanner := bufio.NewScanner(lf)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "[ERROR]") || strings.Contains(scanner.Text(), "[WARNING]") {
			fmt.Println(scanner.Text())
		}
	}
}
