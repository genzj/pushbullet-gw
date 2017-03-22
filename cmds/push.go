package cmds

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

func init() {
	commands = append(commands, cli.Command{
		Name:    "push",
		Usage:   "push to specified devices",
		Aliases: []string{"p"},
		Before:  requireToken,
		Subcommands: []cli.Command{
			cli.Command{
				Name:      "note",
				Usage:     "push a note to specified devices",
				ArgsUsage: "DeviceID Title Body",
				Action: func(c *cli.Context) error {
					if len(c.Args()) < 3 {
						log.Error("DeviceID, Title and Body are required arguments")
						return fmt.Errorf("DeviceID, Title and Body are required arguments")
					} else if resp, err := client.PushNote(c.Args().Get(0), c.Args().Get(1), c.Args().Get(2)); err != nil {
						log.Error(err)
						return err
					} else {
						log.Info(resp)
						prettyPrintStruct(resp)
						return nil
					}
				},
			},
			cli.Command{
				Name:      "link",
				Usage:     "push a link to specified devices",
				ArgsUsage: "DeviceID Title Body Link",
				Action: func(c *cli.Context) error {
					if len(c.Args()) < 4 {
						log.Error("DeviceID, Title, Body and URL are required arguments")
						return fmt.Errorf("DeviceID, Title, Body and URL are required arguments")
					} else if resp, err := client.PushLink(c.Args().Get(0), c.Args().Get(1), c.Args().Get(2), c.Args().Get(3)); err != nil {
						log.Error(err)
						return err
					} else {
						log.Info(resp)
						prettyPrintStruct(resp)
						return nil
					}
				},
			},
		},
	})
}
