package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Tech struct {
	Name string `json:"name"`
}

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func ReadJSON(filename string) map[string]interface{} {

	fileHandler, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening structure")
		os.Exit(2)
	}

	fileData, _ := ioutil.ReadAll(fileHandler)

	var out map[string]interface{}
	json.Unmarshal(fileData, &out)

	return out
}

func envVariable(key string) string {
	// read and parse configurations
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(homedir + directory + "/.env")
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(2)
	}

	val := os.Getenv(key)
	return val
}

func downloadFile(filepath string, fileurl string) (err error) {

	// Create blank file
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.Header.Set("Authorization", String(10))
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(fileurl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}

func setMethod(args []string) string {
	switch args[0] {
	case "regexp":
		if args[1] == "test" {
			return "POST"
		} else if args[1] == "apply" {
			return "PUT"
		}
	case "orch":
		if args[1] == "push" {
			return "PUT"
		} else {
			return "PATCH"
		}
	case "put":
		return "PATCH"
	case "target":
		if args[1] == "delete" {
			return "DELETE"
		} else if args[1] == "create" {
			return "PATCH"
		}
	}

	return "GET"
}

func StringWithCharset(length int, charset string) string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}
