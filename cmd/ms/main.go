package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"measoft_api_cli/pkg/api"
)

const configTemplate = `[common]
s=%s
p=%s
u=%s
w=%s`

func main() {
	app := &cli.App{
		Name: "MeaSoft API CLI",
		Usage: "is an open source tool that enables you to interact with MeaSoft " +
			"services using commands in your command-line shell.",
		Version: "1.0.0",
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "Create encoded config file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "host",
						Value:    "",
						Usage:    "set database host",
						Required: true,
					},
					&cli.StringFlag{
						Name:  "port",
						Value: "3306",
						Usage: "set database port",
					},
					&cli.StringFlag{
						Name:     "username",
						Value:    "",
						Usage:    "set user name",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "password",
						Value:    "",
						Usage:    "set user password",
						Required: true,
					},
					&cli.StringFlag{
						Name:  "out",
						Value: "courier.ini",
						Usage: "output file",
					},
				},
				Action: func(ctx *cli.Context) error {
					config := fmt.Sprintf(
						configTemplate,
						ctx.String("host"),
						ctx.String("port"),
						ctx.String("username"),
						ctx.String("password"),
					)
					encoded_config, err := api.EncodeText(config)
					if err != nil {
						return err
					}

					ioutil.WriteFile(ctx.String("out"), []byte(encoded_config), 0644)

					log.Print("Created.")
					return nil
				},
			},
			{
				Name:  "encode",
				Usage: "Encode config file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "in",
						Value: "courier.ini",
						Usage: "input file",
					},
					&cli.StringFlag{
						Name:  "out",
						Value: "courier_encoded.ini",
						Usage: "output file",
					},
				},
				Action: func(ctx *cli.Context) error {
					blob, err := ioutil.ReadFile(ctx.String("in"))
					if err != nil {
						return err
					}

					config := string(blob)
					encoded_config, err := api.EncodeText(config)
					if err != nil {
						return err
					}

					ioutil.WriteFile(ctx.String("out"), []byte(encoded_config), 0644)
					log.Print("Encoded.")
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
