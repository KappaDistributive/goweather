package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stderr, "", 0)
	args := os.Args
	resp, err := http.Get("https://wttr.in/" + args[1] + "?format=1")
	if err != nil {
		logger.Println(err)
		os.Exit(1)
	} else if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(body))
		defer resp.Body.Close()
	} else {
		logger.Printf("Status code %d", resp.StatusCode)
		os.Exit(1)
	}

}
