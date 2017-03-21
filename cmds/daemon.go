package cmds

import (
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

func init() {
	commands = append(commands, cli.Command{
		Name:    "daemon",
		Usage:   "serve simple HTTP GET APIs",
		Aliases: []string{"d"},
		Action: func(c *cli.Context) error {
			log.WithField("context", c.Command.FullName()).Info("vim-go")
			return nil
		},
	})
}
