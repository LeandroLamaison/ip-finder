package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Boostrap() *cli.App {
	app := cli.NewApp()
	app.Name = "IP Finder"
	app.Usage = "Finds IP address and server name out of an url"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "url",
			Value: "www.google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Finds IP address out of an url",
			Flags:  flags,
			Action: findIP,
		},
		{
			Name:   "server",
			Usage:  "Finds server name out of an url",
			Flags:  flags,
			Action: findServer,
		},
	}
	return app
}

func findIP(c *cli.Context) {
	url := c.String("url")

	ips, err := net.LookupIP(url)

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func findServer(c *cli.Context) {
	url := c.String("url")

	servers, err := net.LookupNS(url)

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
