package server

import (
	"github.com/urfave/cli"
	"golang-skeleton/cmd"
	"os"
)

func Serve() {

	app := cli.NewApp()
	app.Name = "GOLANG SKELETON SYSTEM"

	app.Commands = []cli.Command{
		cmd.ServerCMD,
		cmd.ConsumerCMD,
	}

	_ = app.Run(os.Args)
}
