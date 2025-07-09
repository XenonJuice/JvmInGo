package main

import (
	"fmt"
)

func main() {
	var empty string = ""
	var cmd *Cmd = parseCmd()
	if (*cmd).versionFlag {
		fmt.Println("version 1.0")
	} else if (*cmd).helpFlag || (*cmd).class == empty {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classPath: %s class: %s args: %v \n", (*cmd).cpOption, (*cmd).class, (*cmd).args)
}
