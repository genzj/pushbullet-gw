package cmds

import (
	"encoding/json"
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

	if token := c.String("token"); token != "" {
		client.LoadToken(token)
	} else if code := c.String("code"); code != "" {
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

func prettyPrintStruct(v interface{}) {
	s, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(s))
}

func requireToken(c *cli.Context) error {
	if client.HasToken() {
		return nil
	} else if err := client.RefreshToken(); err != nil {
		log.Error(
			"Cannot get access token, error: ",
			err,
			"Try get code from url then retry: ",
			client.AuthURL(),
		)
		return err
	} else {
		return nil
	}
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
		cli.StringFlag{
			Name:   "token",
			Hidden: true,
			Usage:  "OAuth access token for API access",
		},
	}
	app.Run(os.Args)

	return nil
}
