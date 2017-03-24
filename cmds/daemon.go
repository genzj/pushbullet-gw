package cmds

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/genzj/pushbullet-gw/server"
	"github.com/urfave/cli"
)

func init() {
	commands = append(commands, cli.Command{
		Name:      "daemon",
		Usage:     "serve simple HTTP GET APIs",
		Aliases:   []string{"d"},
		ArgsUsage: "[IP]:Port",
		Action: func(c *cli.Context) error {
			if !c.Args().Present() {
				log.Error("Listen port is mandatory.")
				return fmt.Errorf("Listen IP:port is mandatory, e.g. '127.0.0.1:1323' or ':1323'.")
			}
			server.Start(client, c.Args().Get(0))
			return nil
		},
	})
}
