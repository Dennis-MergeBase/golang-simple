// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package API

// [START import_libraries]
import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
	"log"
	"strings"

	"net/http"
	//"strings"

	"cloud.google.com/go/dialogflow/apiv2"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

// [END import_libraries]
func UpdateAnIntent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	trainingPhrasesParts := strings.Split(params[TrainingPhraseParts], ",")
	messageTexts := strings.Split(params[MessageTexts], ",")
	fmt.Println("display name", params[DisplayName], "trainingPhrasesParts: ", trainingPhrasesParts, ", and ", messageTexts)

	fmt.Printf("Creating intent %s under projects/%s/agent...\n", params[DisplayName], params[ProjectID])
	err := CreateIntent(params[ProjectID], params[ProjectCred], params[DisplayName], trainingPhrasesParts, messageTexts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Done!\n")
}
func UpdateIntent(projectID, permissionsPathJSON, displayName string, trainingPhraseParts, messageTexts []string) error {
	ctx := context.Background()

	intentsClient, clientErr := dialogflow.NewIntentsClient(ctx, option.WithCredentialsFile(permissionsPathJSON))
	if clientErr != nil {
		return clientErr
	}
	defer intentsClient.Close()

	if projectID == "" || displayName == "" {
		return errors.New(fmt.Sprintf("Received empty project (%s) or intent (%s)", projectID, displayName))
	}


	var targetTrainingPhrases []*dialogflowpb.Intent_TrainingPhrase
	var targetTrainingPhraseParts []*dialogflowpb.Intent_TrainingPhrase_Part
	for _, partString := range trainingPhraseParts {
		part := dialogflowpb.Intent_TrainingPhrase_Part{Text: partString}
		targetTrainingPhraseParts = []*dialogflowpb.Intent_TrainingPhrase_Part{&part}
		targetTrainingPhrase := dialogflowpb.Intent_TrainingPhrase{Type: dialogflowpb.Intent_TrainingPhrase_EXAMPLE, Parts: targetTrainingPhraseParts}
		targetTrainingPhrases = append(targetTrainingPhrases, &targetTrainingPhrase)
	}

	intentMessageTexts := dialogflowpb.Intent_Message_Text{Text: messageTexts}
	wrappedIntentMessageTexts := dialogflowpb.Intent_Message_Text_{Text: &intentMessageTexts}
	intentMessage := dialogflowpb.Intent_Message{Message: &wrappedIntentMessageTexts}

	target := dialogflowpb.Intent{DisplayName: displayName, WebhookState: dialogflowpb.Intent_WEBHOOK_STATE_UNSPECIFIED, TrainingPhrases: targetTrainingPhrases, Messages: []*dialogflowpb.Intent_Message{&intentMessage}}

	request := dialogflowpb.UpdateIntentRequest{Intent: &target}

	_, requestErr := intentsClient.UpdateIntent(ctx, &request)


	if requestErr != nil {
		return requestErr
	}

	return nil
}