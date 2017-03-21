package cmds

import (
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

func init() {
	commands = append(commands, cli.Command{
		Name:    "push",
		Usage:   "push to specified devices",
		Aliases: []string{"p"},
		Before:  requireToken,
		Action: func(c *cli.Context) error {
			if resp, err := client.GetUser(); err != nil {
				log.Error(err)
			} else {
				log.Info(resp)
			}
			return nil
		},
	})
}
