package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	hash := flag.Bool("hash", false, "return commit hash")
	flag.Parse()
	if *hash {
		resp, err := http.Get("http://127.0.0.1:8090/debug/info")
		if err != nil {
			log.Fatal()
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Fatal()
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal()
		}

		fmt.Println(string(body))
	}
}
