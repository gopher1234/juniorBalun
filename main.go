package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

var Commit = func() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value
			}
		}
	}

	return ""
}()

func debugInfo(w http.ResponseWriter, req *http.Request) {

	fmt.Println(Commit)
}

func main() {
	http.HandleFunc("/debug/info", debugInfo)

	http.ListenAndServe(":8090", nil)
}
