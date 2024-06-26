package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App{ 
	app := cli.NewApp()
	app.Name = "CLI application"
	app.Usage = "Search IPs and Server Names"
	
	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: 	"ip",
			Usage: 	"Search IP addresses on the internet",
			Flags: 	flags,
			Action: searchIps,
		},
		{
			Name: "servers",
			Usage: "Search for server names on the internet",
			Flags: flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context){
	host := c.String("host")

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("[Public IPs]")

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context){
	host := c.String("host")

	servers, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("[Server Names]")
	for _, server := range servers{
		fmt.Println(server.Host)
	}
}