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
		Action: func(c *cli.Context) error {
			if resp, err := client.Me(); err != nil {
				log.Error(err)
			} else {
				log.Info(resp)
			}
			//if err := client.RefreshToken(); err != nil {
			//log.Error(
			//"Cannot get access token, error: ",
			//err,
			//"Try get code from url: ",
			//client.AuthURL(),
			//)
			//return err
			//}
			return nil
		},
	})
}
