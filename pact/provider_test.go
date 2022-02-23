package pact

import (
	"ModanisaBE/server"
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"os"
	"testing"
)

func TestProvider(t *testing.T) {
	port, _ := utils.GetFreePort()
	go server.CreateServer(port)

	pact := dsl.Pact{
		Host:                     "127.0.0.1",
		Provider:                 "Backend",
		Consumer:                 "Frontend",
		DisableToolValidityCheck: true,
	}

	request := types.VerifyRequest{
		ProviderBaseURL: fmt.Sprintf("http://localhost:%d", port),
		PactURLs: []string{
			"https://og24.pactflow.io/pacts/provider/backend/consumer/frontend/latest",
		},
		BrokerToken: os.Getenv("BROKER_TOKEN"),
	}

	verifyResponses, _ := pact.VerifyProvider(t, request)
	fmt.Println(len(verifyResponses), "pact tests run")
}
