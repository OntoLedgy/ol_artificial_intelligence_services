package testing

import (
	"github.com/OntoLedgy/ol_artificial_intelligence_services/code/services/large_language_model_services"
	"github.com/OntoLedgy/storage_interop_services/code/services/documents/json"
	"testing"
)

func TestChatGPTConnectivity(t *testing.T) {

	authToken := json.ReadJson("configuration.json")

	large_language_model_services.ConnectToChatGPT(authToken["authorisation_token"].(string))

}