package tapsync

import (
	"fmt"
	"net/http"
)

func syncHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is the TAP Sync Service")
}

func TapSyncService() {
    http.HandleFunc("/sync", syncHandler)
}
