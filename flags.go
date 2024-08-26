package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type intelArgs struct {
	JSON            bool
	Loop            bool
	Count           bool
	CDN             bool
	Total           bool
	Limit           bool
	Help            bool
	Internal        bool
	NoLimit         bool
	OutOfScopes     bool
	ReverseCompare  bool
	Tag             string
	ContentType     string
	Server          string
	PublicTarget    string
	ExcludeScope    string
	Status          string
	Title           string
	ExcludeDomain   string
	ExcludeProvider string
	Provider        string
	Date            string
	Method          string
	Compare         string
	Body            string
	BodyFile        string
	ResponseHeaders string
	Technology      string
}

func dropFlags(args []string) []string {

	var argsString string = strings.Join(args, " ")

	if strings.Contains(argsString, "--") {

		var re = regexp.MustCompile(`--(.*)`)
		var arr []string

		rep := re.ReplaceAllString(argsString, ``)
		arr = strings.Split(rep, " ")

		return arr[:len(arr)-1]

	}

	return args

}

func defineIntelArgumentFlags(intelFlags *flag.FlagSet, args *intelArgs) {

	intelFlags.Usage = func() {
		fmt.Println(flagHelp)
	}
	intelFlags.StringVar(&args.ResponseHeaders, "response-headers", "", "filter by response headers")
	intelFlags.StringVar(&args.Technology, "technology", "", "set technology")
	intelFlags.StringVar(&args.Server, "server", "", "filter by server header")
	intelFlags.StringVar(&args.ContentType, "content-type", "", "filter by content-type")
	intelFlags.StringVar(&args.Tag, "tag", "", "filter by watchtower tags")
	intelFlags.StringVar(&args.Provider, "provider", "", "set providers")
	intelFlags.StringVar(&args.Status, "status", "", "match status")
	intelFlags.StringVar(&args.Title, "title", "", "match title")
	intelFlags.StringVar(&args.Date, "date", "", "set date")
	intelFlags.StringVar(&args.ExcludeProvider, "exclude-provider", "", "exclude provider from result")
	intelFlags.StringVar(&args.ExcludeScope, "exclude-scope", "", "exclude scope from result")
	intelFlags.StringVar(&args.ExcludeDomain, "exclude-domain", "", "exclude domain from result")
	intelFlags.StringVar(&args.Method, "method", "", "http request method")
	intelFlags.StringVar(&args.Body, "body", "", "request body")
	intelFlags.StringVar(&args.PublicTarget, "public-target", "", "add public target by name")
	intelFlags.StringVar(&args.BodyFile, "body-file", "", "request body file")
	intelFlags.StringVar(&args.Compare, "compare", "", "compare response")

	intelFlags.BoolVar(&args.OutOfScopes, "oos", false, "remove out of scopes")
	intelFlags.BoolVar(&args.ReverseCompare, "rc", false, "reverse compare")
	intelFlags.BoolVar(&args.JSON, "json", false, "show output as json")
	intelFlags.BoolVar(&args.Count, "count", false, "add count=true")
	intelFlags.BoolVar(&args.CDN, "cdn", false, "add cdn=true")
	intelFlags.BoolVar(&args.Internal, "internal", false, "add internal=true")
	intelFlags.BoolVar(&args.NoLimit, "no-limit", false, "add no_limit=true")
	intelFlags.BoolVar(&args.Total, "total", false, "add total=true")
	intelFlags.BoolVar(&args.Loop, "loop", false, "get all pages")
	intelFlags.BoolVar(&args.Limit, "limit", false, "limit results")
	intelFlags.BoolVar(&args.Help, "help", false, "help of flags")
	intelFlags.BoolVar(&args.Help, "h", false, "help of flags")

}

func showHelp(args []string) {

	if len(args) == 2 {

		switch args[1] {
		case "version":
			fmt.Print(version)
			os.Exit(0)
		case "flags":
			fmt.Print(flagHelp)
			os.Exit(0)
		case "get":
			fmt.Print(getHelp)
			os.Exit(0)
		case "orch":
			fmt.Print(orchHelp)
			os.Exit(0)
		case "regexp":
			fmt.Print(regexHelp)
			os.Exit(0)
		case "put":
			fmt.Print(putHelp)
			os.Exit(0)
		case "target":
			fmt.Print(targetHelp)
			os.Exit(0)
		case "all":
			fmt.Print(getHelp + orchHelp + "\n" + regexHelp + "\n" + putHelp + "\n" + targetHelp)
			os.Exit(0)
		default:
			fmt.Print("Invalid help category")
			os.Exit(0)
		}
	} else {
		fmt.Printf(`## Show details of watchtower flags:
watchtower help flags

## Configuration
watchtower init
watchtower init autocomplete
watchtower help version
watchtower update

## Help of commands
watchtower help get
watchtower help orch 
watchtower help regexp
watchtower help put
watchtower help target

Run 'watchtower help all' to show full help of commands
`)
	}
	os.Exit(0)
}
