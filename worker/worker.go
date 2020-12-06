package worker

import (
	"fmt"
	"github.com/jrallison/go-workers"
	"golang-skeleton/config"
)

func NewWorker(cfg *config.Config) {
	workers.Configure(map[string]string{
		// location of redis instance
		"server": fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		// instance of the database
		"database": "0",
		// number of connections to keep open with redis
		"pool": cfg.Redis.Pool,
		// unique process id for this instance of workers (for proper recovery of inprogress jobs on crash)
		"process": "1",
	})
	return
}
