package main

import (
	"ataps/internal/tapsync"
)

func main() {
    r := tapsync.NewTapSyncService()
    r.Router.Run()
}
