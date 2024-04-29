package main

import (
	"ataps/internal/tapsync"
	"fmt"
	"net/http"
)

func main() {
    fmt.Println("Starting TAP Service")
    tapsync.TapSyncService()
    http.ListenAndServe(":8080", nil)
}
