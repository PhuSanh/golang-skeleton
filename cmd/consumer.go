package cmd

import (
	"github.com/jrallison/go-workers"
	"github.com/urfave/cli"
	"golang-skeleton/config"
	"golang-skeleton/worker"
	"log"
)

var ConsumerCMD = cli.Command{
	Name:    "consumer",
	Aliases: []string{"consumer"},
	Usage:   "Consume queue",
	Action: func(ctx *cli.Context) (err error) {
		// ----- Init config
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatal("cannot load config: ", err)
		}

		worker.NewWorker(cfg)

		// pull messages from "myqueue" with concurrency of 10
		workers.Process("myqueue", myJob, 1)

		// stats will be available at http://localhost:8081/stats
		go workers.StatsServer(8081)

		// Blocks until process is told to exit via unix signal
		workers.Run()

		return
	},
}

func myJob(message *workers.Msg) {
	log.Println("---- myJob message ----")
	log.Println(message)
	// do something with your message
	// message.Jid()
	// message.Args() is a wrapper around go-simplejson (http://godoc.org/github.com/bitly/go-simplejson)
}
