package main

import (
	"ataps/internal/tapsync"
)

func main() {
    r := tapsync.TapSyncService()
    r.Run()
}
