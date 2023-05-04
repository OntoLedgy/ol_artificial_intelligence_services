package testing

import (
	"fmt"
	"github.com/OntoLedgy/ol_artificial_intelligence_services/code/services/large_language_model_services/chatgpt_service"
	"github.com/OntoLedgy/ol_artificial_intelligence_services/code/services/large_language_model_services/chatgpt_service/prompts"
	"github.com/OntoLedgy/storage_interop_services/code/services/documents/json"
	"testing"
)

func TestChatGPTConnectivity(t *testing.T) {

	authToken := json.ReadJson("configuration.json")

	client := chatgpt_service.ConnectToChatGPT(authToken["authorisation_token"].(string))

	requestMessage := prompts.CleanCodingInitialiser

	response := chatgpt_service.CreateChatCompletion(
		client,
		requestMessage)

	fmt.Println(response)

}
