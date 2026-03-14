package main

import (
	"fmt"
	"log"
	"os"
	"github.com/urfave/cli/v2"
)

func main(){
	//it used to create cli application object using the package
	app:=&cli.App{
		//they are field of structure
		Name:"Health Checker",
		Usage:"A tiny tool that checks wheather a website running or is down",
		//cli.StringFlag is a struct provided by the CLI library to define a string command-line flag.
		Flags:[]cli.Flag{
			&cli.StringFlag{
				Name: "domain",
				Aliases: []string{"d"},
				Usage: "Domain name to check",
				Required: true,
			},
			&cli.StringFlag{
				Name:"port",
				Aliases: []string{"p"},
				Usage: "Port number to check",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error{
			port := c.String("port")
			if c.String("port")==""{
				port="80"
			}
			status:=Check(c.String("domain"),port)
			fmt.Println(status)
			return nil
		},
	}	
	err:=app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}