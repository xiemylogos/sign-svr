package main

import (
	"github.com/urfave/cli"
	"github.com/xiemylogos/sign-svr/cmd"
	"os"
	"runtime"
)

func setupAPP() *cli.App {
	app := cli.NewApp()
	app.Usage = "sign service"
	app.Action = startAgent
	app.Commands = []cli.Command{
		cmd.SignCommand,
		cmd.SignHashCommand,
	}
	app.Before = func(context *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}
	return app
}
func main() {
	if err := setupAPP().Run(os.Args); err != nil {
		os.Exit(1)
	}
}

func startAgent(ctx *cli.Context) {
}
