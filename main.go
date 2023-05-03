package main

import (
	"fmt"
	"github.com/OntoLedgy/storage_interop_services/code/services/documents/json"
)

var (
	// Version These fields are populated by govvv
	Version    string
	BuildDate  string
	GitCommit  string
	GitBranch  string
	GitState   string
	GitSummary string
)

func main() {

	fmt.Printf("Starting AI Services Version %s-%s-%s\n", Version, GitCommit, BuildDate)
	json.ReadJson("test")

	//large_language_model_services.ConnectToChatGPT(authToken["authorisation_token"].(string))

}
