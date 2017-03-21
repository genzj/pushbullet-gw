package cmds

import (
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

func init() {
	commands = append(commands, cli.Command{
		Name:    "device",
		Usage:   "maintain registered devices",
		Aliases: []string{"p"},
		Before:  requireToken,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "list",
				Usage: "show all registered devices",
			},
		},
		Action: deviceAction,
	})
}

func deviceAction(c *cli.Context) error {
	list := c.Bool("list")
	if list {
		if resp, err := client.ListDevices(); err != nil {
			log.Error(err)
		} else {
			prettyPrintStruct(resp)
		}
	}

	return nil
}
