package API

import (
	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"context"
	"fmt"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	"net/http"
)


func ExportAgentToZip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	ctx := context.Background()

	intentsClient, clientErr := dialogflow.NewAgentsClient(ctx)
	if clientErr != nil {
		fmt.Fprintf(w, "Problem getting client")
	}
	defer intentsClient.Close()

	request := &dialogflowpb.ExportAgentRequest{Parent: "projects/createagents"}

	op, err := intentsClient.ExportAgent(ctx, request)

	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
	fmt.Println("and my resp is: ", resp.Agent)

}
