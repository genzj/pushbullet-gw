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
				Usage: "show registered devices, output will be filtered by the DeviceID if specified",
			},
		},
		ArgsUsage: "[DeviceID]",
		Action:    deviceAction,
	})
}

func deviceAction(c *cli.Context) error {
	list := c.Bool("list")
	if list {
		if resp, err := client.ListDevices(); err != nil {
			log.Error(err)
		} else if c.Args().Present() {
			deviceID := c.Args().First()
			for _, device := range resp.Devices {
				if device.ID == deviceID {
					prettyPrintStruct(device)
				}
			}
		} else {
			prettyPrintStruct(resp)
		}
	}

	return nil
}
