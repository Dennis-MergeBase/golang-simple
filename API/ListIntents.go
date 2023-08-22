package API

import (
	. "CreateConversationBackend_NLP/ResponseJSON"
	. "CreateConversationBackend_NLP/ErrorHandling"
	"errors"

	"cloud.google.com/go/dialogflow/apiv2"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	"net/http"
	"net/url"
)

func ListAllIntents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	fmt.Println("ProjectID: ",  params[ProjectID], "  ProjectCred: ", params[ProjectCred])

	// TODO: sessionID

	var err error
	var intents []*dialogflowpb.Intent

	credentialsPath, err := url.QueryUnescape(params[ProjectCred])

	if err != nil{
		fmt.Println("problem unescaping credentials")
		ReturnGeneralError(w, http.StatusInternalServerError, err.Error())
		return
	}


	//intents, err = ListIntents(params[ProjectID], "/Users/vee/Desktop/" + credentialsPath)
	intents, err = ListIntents(params[ProjectID], "/" + credentialsPath)


	//
	if err != nil {
		ReturnGeneralError(w, http.StatusBadRequest, err.Error())
		return
	}

	//fmt.Fprintf(w,"Number of intents: %d\n", len(intents))
	if len(intents) > 0 {
		w.WriteHeader(http.StatusOK)
		jData, _ := json.Marshal(ReturnResponse{http.StatusOK, intents})
		_, err := w.Write(jData)
		if err != nil{
			ReturnGeneralError(w, http.StatusInternalServerError, err.Error())
			return
		}
		return

	}else{
		ReturnGeneralError(w, http.StatusNotFound, err.Error())
		return
	}
}

// [START dialogflow_list_intents]

func ListIntents(projectID string, permissionsJSONPath string) ([]*dialogflowpb.Intent, error) {
	fmt.Println("projectID: ", projectID)
	fmt.Println("permissionsPath: ", permissionsJSONPath)

	ctx := context.Background()

	intentsClient, clientErr := dialogflow.NewIntentsClient(ctx, option.WithCredentialsFile(permissionsJSONPath))
	if clientErr != nil {
		return nil, clientErr
	}
	defer intentsClient.Close()

	if projectID == "" {
		return nil, errors.New(fmt.Sprintf("Received empty project (%s)", projectID))
	}

	parent := fmt.Sprintf("projects/%s/agent", projectID)

	request := dialogflowpb.ListIntentsRequest{Parent: parent}

	intentIterator := intentsClient.ListIntents(ctx, &request)

	var intents []*dialogflowpb.Intent


	for intent, status := intentIterator.Next(); status != iterator.Done; {
		if intent != nil{
			intents = append(intents, intent)
			intent, status = intentIterator.Next()
		}else{
			return nil, status
		}
	}

	return intents, nil
}


// [END dialogflow_list_intents]
