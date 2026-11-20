package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mossman1215/metaname-cli/internal"
)

/*
*
handle cli execution
argument parsing
environment variables
sub command resolution
error handling
help text
*/
var identifier *string
var apiToken *string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	supportedMethods := []string{
		"create_dns_record",
		"dns_zone",
		"domain_name",
		"update_dns_record",
		"update_name_servers",
		"delete_dns_record",
	}
	command := flag.String("command", "", "command value known in metaname as method")
	jsonFile := flag.String("jsonFile", "", "Path to a json file for raw uploads")
	test := flag.Bool("test", false, "run command against test api server")
	flag.Parse()
	commandLower := strings.ToLower(*command)

	switch commandLower {
	case "raw":
		if *jsonFile == "" {
			fmt.Println("jsonFile argument not found")
			os.Exit(1)
		}
		jsonFilePath, err := filepath.Abs(*jsonFile)
		check(err)
		dat, err := os.ReadFile(jsonFilePath)
		check(err)
		rawUpload(*test, string(dat))
	default:
		fmt.Printf("command %s not supported\n", *command)
		fmt.Printf("supported commands\n raw\n")
		fmt.Println(supportedMethods)
		os.Exit(1)
	}
}

/**
* take command raw and a filepath to a json payload
* check the file is json
* upload to api via requests.go
 */
func rawUpload(test bool, jsonPayload string) {
	internal.MetanameRequest(test, jsonPayload)
}
