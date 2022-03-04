package main

import (
	"ip-finder/app"
	"log"
	"os"
)

func main() {
	ipFinder := app.Boostrap()

	if err := ipFinder.Run(os.Args); err != nil {
		log.Fatal("Unexpeced error running IP Finder")
	}
}
