package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/abhijithda/robotic-automation-library/robot"
)

func main() {
	robot.RegisterCmdOptions()
	flag.Parse()
	if robot.CmdOptions.LogFile != "" {
		LogFile, err := os.OpenFile(robot.CmdOptions.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file %s: %v", robot.CmdOptions.LogFile, err)
		}
		defer LogFile.Close()
		log.SetOutput(LogFile)

		fmt.Printf("Log file: %s\n\n", robot.CmdOptions.LogFile)
	}

	stack := robot.Sort(robot.CmdOptions.Width, robot.CmdOptions.Height, robot.CmdOptions.Length, robot.CmdOptions.Mass)
	fmt.Printf("The package should go to the %s stack.\n\n", stack)
}
