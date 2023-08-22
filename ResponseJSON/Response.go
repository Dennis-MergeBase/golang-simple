package ResponseJSON

import (
	. "CreateConversationBackend_NLP/ErrorHandling"
	"encoding/json"
	"fmt"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	"io"
	"net/http"
)


type IntentResponse struct{
	name string
	displayName string
	action string
	followupIntent string
	parentFollowupIntent string
	inputContextNames []string
	messages []string
}

type ReturnResponse struct{
	Status   int
	Response []*dialogflowpb.Intent
}

type ReturnErrorResponse struct{
	Status int
	Response string
}

type ReturnSuccessResponse struct{
	Status int
	Response string
}

func ReturnCreateIntentSuccess(w http.ResponseWriter, r http.Request){
	returnMessage, err := json.Marshal(ReturnSuccessResponse{http.StatusOK, "Your intent was successfully created."})
	if err != nil {
		fmt.Println(err)
		ReturnCreateIntentError(w, &r)
		return
	}
	io.WriteString(w, string(returnMessage))
}


func CreateReturnResponse(response []*dialogflowpb.Intent) *ReturnResponse {
	return &ReturnResponse{Status: http.StatusOK, Response: response}
}

func CreateIntentResponse(name, displayName,action, followupIntent, parentFollowupIntent string, inputContextNames, messages []string) *IntentResponse {
	return &IntentResponse{name, displayName, action, followupIntent, parentFollowupIntent, inputContextNames, messages}
}

