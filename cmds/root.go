package cmds

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/genzj/pushbullet-gw/pushbullet"
	"github.com/go-resty/resty"
	"github.com/urfave/cli"
)

var app cli.App
var commands cli.Commands

var client *pushbullet.Client

func createClient(c *cli.Context) error {
	for c.Parent() != nil {
		c = c.Parent()
	}

	clientID := c.String("client-id")
	clientSecret := c.String("client-secret")
	redirectURI := c.String("redirect-uri")

	if clientID == "" || clientSecret == "" || redirectURI == "" {
		log.Errorf("ERROR: client-id, client-secret and redirect-uri cannot be blank")

		return fmt.Errorf("ERROR: client-id, client-secret and redirect-uri cannot be blank")
	}

	client = pushbullet.NewClient(clientID, clientSecret, redirectURI)

	if code := c.String("code"); code != "" {
		client.Credential.Code = code
	}
	return nil
}

func setLogger(c *cli.Context) error {
	debug := c.Bool("debug")
	resty.SetDebug(debug)
	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	return nil
}

func Run() error {
	app := cli.NewApp()
	app.Usage = "access Pushbullet APIs by commands and simple HTTP GET requests"
	app.Version = "0.1.0"
	app.Author = "genzj <zj0512@gmail.com>"
	app.Commands = commands
	app.Before = func(c *cli.Context) error {
		fs := [](func(c *cli.Context) error){
			setLogger,
			createClient,
		}
		for _, f := range fs {
			if err := f(c); err != nil {
				return err
			}
		}
		return nil
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "show more debug messages",
		},
		cli.StringFlag{
			Name:  "client-id",
			Usage: "Client ID of the OAuth application",
		},
		cli.StringFlag{
			Name:  "client-secret",
			Usage: "Client secret of the OAuth application",
		},
		cli.StringFlag{
			Name:  "code",
			Usage: "OAuth code used to get access token",
		},
		cli.StringFlag{
			Name:  "redirect-uri",
			Value: "http://127.0.0.1/code",
			Usage: "URI to receive authentication code of OAuth",
		},
	}
	app.Run(os.Args)

	return nil
}
