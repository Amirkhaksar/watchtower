package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"sync"
)

func initAutoComplete(watch_directory string) {

	err := downloadFile(watch_directory+"/_watchtower", repository+"/main/_watchtower")
	if err != nil {
		fmt.Println(err)
	}

	err = downloadFile(watch_directory+"/init-autocomplete.sh", repository+"/main/init-autocomplete.sh")
	if err != nil {
		fmt.Println(err)
	}

	cmd := exec.Command("zsh", watch_directory+"/init-autocomplete.sh")
	_, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

}

func updateWatchtowerFiles(watch_directory string) {
	err := downloadFile(watch_directory+"/structure.json", repository+"/main/structure.json")
	if err != nil {
		fmt.Println(err)
	}

	err = downloadFile(watch_directory+"/_watchtower", repository+"/main/_watchtower")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("'structure.json' and '_watchtower' updated!")

}

func initWatchtower(watch_directory string) {

	err := downloadFile(watch_directory+"/.env", repository+"/main/.env")
	if err != nil {
		fmt.Println(err)
	}

	err = downloadFile(watch_directory+"/structure.json", repository+"/main/structure.json")
	if err != nil {
		fmt.Println(err)
	}

}

func setPublicTargetBody(target_name string) string {
	return `{ "name": "` + target_name + `", "dns_brute": { "type": [ "static", "dynamic" ], "interval": 0 }, "eligible_for_bounty": true, "domains": [], "filter": { "firewall": [], "regex": [] }, "flags": [ "cdn", "tech_detect" ], "http_options": { "type": "new", "": 10, request_per_second"ports": [ 443 ], "retry_on_failure": 0, "deny_statuses": [ "400", "5xx" ], "deny_cdn": true }, "name_resolution": { "type": "all", "resolvers": [ "8.8.4.4", "129.250.35.251", "208.67.222.222" ] }, "on_start": { "enumeration": false, "dns_brute": false, "name_resolution": false }, "out_of_scopes": [], "providers": { "enumeration": [ "crtsh", "subfinder", "abuseipdb", "chaos", "dnscrawl", "sourcegraph" ] }, "source": "other" }`
}

func setAPI(api string, flagArgs intelArgs) string {

	if flagArgs.Count {
		api = fmt.Sprintf("%s&count=true", api)
	}
	if flagArgs.CDN {
		api = fmt.Sprintf("%s&cdn=true", api)
	}
	if flagArgs.Internal {
		api = fmt.Sprintf("%s&internal=true", api)
	}
	if flagArgs.NoLimit {
		api = fmt.Sprintf("%s&no_limit=true", api)
		flagArgs.Loop = false
	}
	if flagArgs.Total {
		api = fmt.Sprintf("%s&total=true", api)
	}
	if flagArgs.JSON {
		api = fmt.Sprintf("%s&json=true", api)
		flagArgs.Loop = false
	}
	if flagArgs.Provider != "" {
		api = fmt.Sprintf("%s&provider=%s", api, URLEncode(flagArgs.Provider))
	}
	if flagArgs.Title != "" {
		api = fmt.Sprintf("%s&title=%s", api, URLEncode(flagArgs.Title))
	}
	if flagArgs.Status != "" {
		api = fmt.Sprintf("%s&status=%s", api, URLEncode(flagArgs.Status))
	}
	if flagArgs.Date != "" {
		api = fmt.Sprintf("%s&date=%s", api, URLEncode(flagArgs.Date))
	}
	if flagArgs.ExcludeDomain != "" {
		api = fmt.Sprintf("%s&exclude_domain=%s", api, URLEncode(flagArgs.ExcludeDomain))
	}
	if flagArgs.ExcludeScope != "" {
		api = fmt.Sprintf("%s&exclude_scope=%s", api, URLEncode(flagArgs.ExcludeScope))
	}
	if flagArgs.ExcludeProvider != "" {
		api = fmt.Sprintf("%s&exclude_provider=%s", api, URLEncode(flagArgs.ExcludeProvider))
	}
	if flagArgs.Tag != "" {
		api = fmt.Sprintf("%s&tag=%s", api, URLEncode(flagArgs.Tag))
	}
	if flagArgs.Technology != "" {
		api = fmt.Sprintf("%s&technology=%s", api, URLEncode(flagArgs.Technology))
	}
	if flagArgs.Limit {
		flagArgs.Loop = false
	}
	if flagArgs.ResponseHeaders != "" {
		headers := parseHeaders(flagArgs.ResponseHeaders)
		for header, value := range headers {
			api = fmt.Sprintf("%s&response_headers[%s]=%s", api, URLEncode(header), URLEncode(value))
		}
	}
	if flagArgs.ContentType != "" {
		api = fmt.Sprintf("%s&response_headers[content_type]=%s", api, URLEncode(flagArgs.ContentType))
	}
	if flagArgs.Server != "" {
		api = fmt.Sprintf("%s&response_headers[server]=%s", api, URLEncode(flagArgs.Server))
	}

	return api
}

func initCompare(response string, flagArgs intelArgs) {
	f, err := os.Create("/tmp/watchtower_client_1")
	if err != nil {
		log.Fatal(err)
	}

	_, err2 := f.WriteString(response)
	if err2 != nil {
		log.Fatal(err2)
	}
	f.Close()

	var f1, f2 string
	f1 = flagArgs.Compare
	f2 = "/tmp/watchtower_client_1"

	if flagArgs.ReverseCompare {
		f1 = "/tmp/watchtower_client_1"
		f2 = flagArgs.Compare
	}

	if _, err := os.Stat(flagArgs.Compare); err != nil {
		fmt.Println("Compare file does not exist!")
		return
	}

	cmd := exec.Command("bash", "-c", "comm -23 <(cat "+f1+"|sort -u) <(cat "+f2+"|sort -u)")
	stdout, _ := cmd.Output()

	fmt.Print(string(stdout))
}

func makeLoop(api string, flagArgs intelArgs, body string, limit string) string {
	var out string

	count, err := strconv.Atoi(MakeHttpRequest(api+"&count=true", flagArgs, body))
	if err != nil {
		fmt.Print(err)
		fmt.Println("an error occurred in sending request to counting data")
		os.Exit(0)
	}

	// Multithreading
	var wg sync.WaitGroup
	results := make(chan string, ((count / 1000) + 1))

	// Fetch results
	wg.Add(((count / 1000) + 1))

	for i := 0; i <= (count / 1000); i++ {

		go func(i int) {
			defer wg.Done()
			results <- MakeHttpRequest(api+"&limit="+limit+"&page="+strconv.Itoa(i), flagArgs, body)
		}(i)

	}

	wg.Wait()
	close(results)

	i := 0
	for result := range results {

		if flagArgs.Compare == "" {

			if i == (count / 1000) {
				fmt.Print(result)
			} else {
				fmt.Print(result + "\n")
			}

		} else {

			if i == (count / 1000) {
				out += result
			} else {
				out += result + "\n"
			}
		}
		i++
	}

	return out
}

func parseTechnologies(resp string, args []string) {
	if args[0] == "get" && args[1] == "technologies" && args[2] == "list" {
		var items []Tech

		// Parse the JSON data
		err := json.Unmarshal([]byte(resp), &items)
		if err != nil {
			log.Fatalf("Error parsing JSON: %v", err)
		}

		// Iterate over the items and print the name field
		for _, item := range items {
			fmt.Println(item.Name)
		}
		os.Exit(1)
	}
}
