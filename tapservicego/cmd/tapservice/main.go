package main

import (
	"ataps/internal/tapsync"
	"fmt"
)

func main() {
	config := tapsync.NewConfig() // by default config uses env variables
	r := tapsync.NewTapSyncService(config)
	r.Router.Run(fmt.Sprintf(":%d", config.Port))
	r.DB.Close()
}
