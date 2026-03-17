package main

import (
	check "Health_Checkup/logic"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	//it used to create cli application object using the package
	app := &cli.App{
		//they are field of structure
		Name:  "Health Checker",
		Usage: "A tiny tool that checks wheather a website running or is down",
		//cli.StringFlag is a struct provided by the CLI library to define a string command-line flag.
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "domain",
				Aliases: []string{"d"},
				Usage:   "Domain name to check",
				//Required: true,
			},
			&cli.StringFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "Port number to check",
				//Required: false,
			},
		},
		//Action is the function that runs cli commands
		//cli.Context is a structure (struct) provided by the urfave/cli package.It stores information about the CLI command execution, such as:
		//command arguments,flags,command name,app info
		//simple word - cli.Context stores command data
		Action: func(c *cli.Context) error {
			//get the port value
			port := c.String("port")
			//If the user did not provide a port, then use port 80.
			if c.String("port") == "" {
				port = "80"
			}
			status := check.Check(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}
	//app.Run - this start my cli application
	//err store any err
	err := app.Run(os.Args)
	if err != nil {
		//print the err and stop imidiatelly
		log.Fatal(err)
	}
}
