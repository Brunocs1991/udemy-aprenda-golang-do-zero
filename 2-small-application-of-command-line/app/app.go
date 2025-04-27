package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar cria uma nova instância de cli.App com configurações padrão
// e retorna o ponteiro para essa instância.
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Application of command line"
	app.Usage = "Find ips and names of servers in the network"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Find the ip of a server in the network",
			Flags:  flags,
			Action: findIp,
		},
		{
			Name:   "servers",
			Usage:  "Find the name of a server in the network",
			Flags:  flags,
			Action: findServers,
		},
	}
	return app
}

func findIp(c *cli.Context) {
	host := c.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func findServers(c *cli.Context) {
	host := c.String("host")
	names, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, name := range names {
		fmt.Println(name.Host)
	}
}
