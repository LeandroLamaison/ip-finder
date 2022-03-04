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
	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Finds IP address out of an url",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "url",
					Value: "www.google.com.br",
				},
			},
			Action: findIP,
		},
	}
	return app
}

func findIP(c *cli.Context) {
	url := c.String("url")

	ips, err := net.LookupIP(url)

	if err != nil {
		log.Fatal("URL coudn't be reached.")
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
