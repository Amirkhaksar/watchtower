package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"
)

func createDirectory(directory string) {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		if err := os.Mkdir(directory+"", os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}

func defaultLoop(args []string) bool {
	if args[0] == "get" {
		if args[1] == "lives" || args[1] == "fresh" || args[1] == "subdomains" || args[1] == "latest" || args[1] == "http" {
			return true
		}
	}
	return false
}

func readBody(flagArgs intelArgs) string {

	if flagArgs.Body != "" {

		return flagArgs.Body

	} else if flagArgs.BodyFile != "" {

		// read body file
		fileContent, err := ioutil.ReadFile(flagArgs.BodyFile)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(0)
		}

		return string(fileContent)
	}

	return ""
}

func parseAPI(api string, args []string) string {
	api = strings.Replace(api, "{{arg}}", args[len(args)-1], -1)
	api = strings.Replace(api, "{{base}}", envVariable("baseURL"), -1)
	return api
}

func URLEncode(s string) string {
	return url.QueryEscape(s)
}

func parseHeaders(headerStr string) map[string]string {
	headers := make(map[string]string)
	headerPairs := strings.Split(headerStr, ";")
	for _, pair := range headerPairs {
		parts := strings.SplitN(strings.TrimSpace(pair), ":", 2)
		if len(parts) == 2 {
			headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}
	return headers
}
