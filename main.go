package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	repository = "https://raw.githubusercontent.com/Mr-MSA/watchtower"
	version    = "1.2.0"
	directory  = "/.watchtower"
)

func main() {

	// Read home directoru location
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	// Set watchtower local directory path
	watch_directory := homedir + directory

	// Drop flags from arguments
	args := dropFlags(os.Args[1:])
	// Check length of args to show help text
	if len(args) == 0 {
		fmt.Println("Help of commands: watchtower help")
		os.Exit(0)
	}

	switch args[0] {
	case "init":
		createDirectory(watch_directory)

		if len(args) == 2 && args[1] == "autocomplete" {

			initAutoComplete(watch_directory)

			os.Exit(0)
		}

		initWatchtower(watch_directory)

		os.Exit(0)
	case "update":

		updateWatchtowerFiles(watch_directory)
		os.Exit(0)
	}

	// check config directory exists
	if _, err := os.Stat(watch_directory); os.IsNotExist(err) {
		fmt.Println("Directory " + watch_directory + " doesn't exist! please execute 'watchtower init' ")
		os.Exit(0)
	}

	// validate baseurl
	if envVariable("baseURL") == "WATCH_SERVER" {
		fmt.Println("Please set watchtower server address at " + watch_directory + "/.env")
		os.Exit(0)
	}

	// read and parse configurations
	config := ReadJSON(watch_directory + "/structure.json")

	// show help
	if args[0] == "help" {
		showHelp(args)
	}

	// set variables
	var api string
	var count = 1
	var limit string

	if envVariable("resultsLimit") != "" {
		limit = envVariable("resultsLimit")
	} else {
		limit = "1000"
	}

	// find endpoint
	var conf map[string]interface{} = config
	for i := 0; i <= len(args)-1; i++ {

		if conf[args[i]] == nil {
			break
		}

		if fmt.Sprintf("%T", conf[args[i]]) == "string" {

			api = conf[args[i]].(string)
			count += i
			break
		} else {

			conf = conf[args[i]].(map[string]interface{})
		}
	}

	// check the api string
	if api == "" {
		fmt.Println("Command/API not found")
		os.Exit(0)
	}

	// check the command has required to be input
	if strings.Contains(api, "{{arg}}") {
		count++
	}

	if len(os.Args[1:]) < count {
		fmt.Println("Command/API not found")
		os.Exit(0)
	}

	// parse api endpoint
	api = parseAPI(api, args)

	// parse flags
	var flagArgs intelArgs
	intelCommand := flag.NewFlagSet("main", flag.ContinueOnError)
	defineIntelArgumentFlags(intelCommand, &flagArgs)

	if err := intelCommand.Parse(os.Args[1:][count:]); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(2)
	}

	// help
	if flagArgs.Help {
		fmt.Printf("%s", flagHelp)
		os.Exit(3)
	}

	// set body
	body := readBody(flagArgs)

	if flagArgs.PublicTarget != "" {
		body = setPublicTargetBody(flagArgs.PublicTarget)
	}

	// append ?
	if !strings.Contains(api, "?") {
		api = api + "?"
	}

	// set default request methods
	if flagArgs.Method == "" {
		flagArgs.Method = setMethod(args)
	}

	// set default loop
	if defaultLoop(args) {
		flagArgs.Loop = true
	}

	// set endpoint by flags
	api = setAPI(api, flagArgs)

	var out string
	if flagArgs.Loop && !flagArgs.Count && !flagArgs.NoLimit {

		out = makeLoop(api, flagArgs, body, limit)

	} else {

		// send http request to api endpoint
		resp := MakeHttpRequest(api, flagArgs, body)

		// return techs list as json
		parseTechnologies(resp, args)

		// check if compare enable save response to a variable
		if flagArgs.Compare == "" {
			fmt.Print(resp)
		} else {
			out = resp
		}

	}

	if flagArgs.Compare != "" {
		initCompare(out, flagArgs)
	}

}
