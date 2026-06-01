package main

import (
	"JvmInGo/classpath"
	"fmt"
	"strings"
)

const dot = "."
const slash = "/"

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
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classPath: %s class: %s args: %v \n", cp, (*cmd).class, (*cmd).args)
	className := strings.Replace(cmd.class, dot, slash, -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Can not find or load main class %s\n", className)
		return
	}
	fmt.Printf("class data : %v\n", classData)
}
