package main

import (
	"ead-tools/cmds"
	"fmt"
	"os"
)

func main() {
	//fmt.Print("== ead-tools")
	if err := cmds.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error(), "\n")
		cmds.CloseLog()
		cmds.ScanLog()
		os.Exit(1)
	}

}
