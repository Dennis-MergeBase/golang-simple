package API

import (
	"cloud.google.com/go/dialogflow/apiv2"
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	"net/http"
	response "CreateConversationBackend_NLP/ResponseJSON"
	errorHandling "CreateConversationBackend_NLP/ErrorHandling"
)

func DeleteAnIntent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	if params[IntentPath] == "" {
		errorHandling.ReturnGeneralError(w, http.StatusPartialContent,"Expected non-empty -intention-id argument\n")
	}

	// fmt.Printf("Deleting intent projects/%s/agent/intents/%s...\n", params[ProjectID], params[IntentPath])
	err := DeleteIntent( params[ProjectID], params[ProjectCred], params[IntentPath] )
	if err != nil {
		errorHandling.ReturnGeneralError(w, http.StatusInternalServerError, err.Error())
	}
	// fmt.Fprintf(w, "Intent has been successfully deleted!\n")
	response.ReturnCreateIntentSuccess(w, *r)
}

// [START dialogflow_delete_intent]
func DeleteIntent(projectID, permissionsJSONPath, intentID string) error {
	ctx := context.Background()

	intentsClient, clientErr := dialogflow.NewIntentsClient(ctx, option.WithCredentialsFile(permissionsJSONPath))

	if clientErr != nil {
		return clientErr
	}
	defer intentsClient.Close()

	if projectID == "" || intentID == "" {
		return errors.New(fmt.Sprintf("Received empty project (%s) or intent (%s)", projectID, intentID))
	}

	targetPath := fmt.Sprintf("projects/%s/agent/intents/%s", projectID, intentID)

	request := dialogflowpb.DeleteIntentRequest{Name: targetPath}

	requestErr := intentsClient.DeleteIntent(ctx, &request)
	if requestErr != nil {
		return requestErr
	}

	return nil
}